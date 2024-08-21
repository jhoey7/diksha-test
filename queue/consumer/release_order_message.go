package consumer

import (
	"edot-test/queue/initialize"
	"edot-test/repositories"
	"edot-test/services"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/streadway/amqp"
	"log/slog"
)

// ReleaseOrderQueueName queue name
var (
	ReleaseOrderQueueName = beego.AppConfig.DefaultString("release.order.queue.name", "release_order_queue")
)

// StartReleaseOrderConsumer function to consume release order
func StartReleaseOrderConsumer(threadName string) *amqp.Channel {
	slog.Info("startCancelledPurchaseConsumer ...", threadName)
	ch, err := initialize.GetRabbitConn().Channel()
	initialize.CFailOnError(err, "Failed to connect channel")

	q, err := ch.QueueDeclare(
		ReleaseOrderQueueName, // name
		true,                  // durable
		false,                 // delete when unused
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // arguments
	)

	initialize.CFailOnError(err, "Failed to declare a queue")

	deliveries, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	initialize.CFailOnError(err, "Failed to register a consumer")

	go sendReleaseOrderMessageHandle(deliveries, threadName)
	return ch
}

// sendReleaseOrderMessageHandle function to handle consume release order message.
func sendReleaseOrderMessageHandle(deliveries <-chan amqp.Delivery, threadName string) {
	slog.Info("sendCancelledPurchaseMessageHandle ...", threadName)
	for d := range deliveries {
		slog.Info(
			"[%v]got %dB delivery: [%v] %q",
			threadName,
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		o := orm.NewOrm()
		productRepo := repositories.NewProductRepository(o)
		productDtl := repositories.NewProductDetailRepository(o)
		orderRepo := repositories.NewOrderRepository(o)
		orderDtlRepo := repositories.NewOrderDetailRepository(o)

		svc := services.NewCheckoutOrderService(productRepo, productDtl, orderRepo, orderDtlRepo, o, 1)
		svc.ReleaseOrder(d.Body)

		d.Ack(false)
	}
}

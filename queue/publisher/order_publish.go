package publisher

import (
	"edot-test/models"
	"edot-test/queue/initialize"
	"github.com/astaxie/beego"
	"log/slog"

	"encoding/json"
	"github.com/streadway/amqp"
)

// ReleaseOrderQueueName queue name
var (
	ReleaseOrderQueueName        = beego.AppConfig.DefaultString("release.order.queue.name", "release_order_queue")
	ReleaseOrderDelayedQueueName = beego.AppConfig.DefaultString("release.order.queue.delayed.name", "release_order_queue_delayed")
)

// ReleaseOrderMessage function to declare and publish queue for release qty.
func ReleaseOrderMessage(message models.OrderReleaseQueue, delayTimeout int) {
	data, err := json.Marshal(message)

	ch, err := initialize.GetRabbitConn().Channel()
	initialize.CFailOnError(err, "Failed to connect channel")
	defer ch.Close()

	args := make(amqp.Table)
	args["x-dead-letter-exchange"] = ""
	args["x-dead-letter-routing-key"] = ReleaseOrderQueueName
	args["x-message-ttl"] = delayTimeout * 60000 // minutes

	q, err := ch.QueueDeclare(
		ReleaseOrderDelayedQueueName, // name
		true,                         // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		args,                         // arguments
	)
	initialize.CFailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         data,
		})
	initialize.CFailOnError(err, "Failed to publish a queue")
	slog.Info("Publish ReleaseOrderMessage", message, err)
}

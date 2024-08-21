package initialize

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/streadway/amqp"
	"log"
)

var (
	conn *amqp.Connection
	err  error
)

// InitRabbitConn : factory method to initialize rabbit connection
func InitRabbitConn() {
	var cErr error
	conn, cErr = amqp.Dial("amqp://" + beego.AppConfig.String("userRabbit") + ":" + beego.AppConfig.String("passRabbit") + "@" + beego.AppConfig.String("urlRabbit"))
	failOnError(cErr, "Failed to connect to RabbitMQ")
}

// GetRabbitConn : get rabbit connection
func GetRabbitConn() *amqp.Connection {
	if conn == nil {
		conn, err = amqp.Dial("amqp://" + beego.AppConfig.String("userRabbit") + ":" + beego.AppConfig.String("passRabbit") + "@" + beego.AppConfig.String("urlRabbit"))
		return nil
	}
	return conn
}

// CFailOnError is function to return error
func CFailOnError(cErr error, msg string) {
	if cErr != nil {
		log.Fatalf("%s: %s", msg, cErr)
	}
}

// failOnError is function
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// PutToQueue : publish message function
func PutToQueue(channel *amqp.Channel, q amqp.Queue, msg []byte) error {
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         msg,
	})

	return err
}

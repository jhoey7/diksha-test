package consumer

import (
	"errors"
	"github.com/streadway/amqp"
)

var allOpenChannels []*amqp.Channel

// StartConsumingQueue start all queue message consumer
func StartConsumingQueue() error {
	if allOpenChannels != nil {
		return errors.New("all queue consumer is already started")
	}

	var openChannel *amqp.Channel

	openChannel = StartReleaseOrderConsumer("ReleaseOrderConsumer")
	allOpenChannels = append(allOpenChannels, openChannel)

	return nil
}

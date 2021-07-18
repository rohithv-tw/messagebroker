package Publisher

import (
	"github.com/google/uuid"
	"log"
	"message-broker/messagebroker/Broker"
	"message-broker/messagebroker/Message"
)

type publisher struct {
	channel string
	message Message.IMessage
	broker  Broker.IBroker
	id      uuid.UUID
}

func Create(broker Broker.IBroker) IPublisher {
	return publisher{
		broker: broker,
		id:     uuid.New(),
	}
}

func (pub publisher) SetContext(channel string, message Message.IMessage) IPublisher {
	pub.message = message
	pub.channel = channel
	return pub
}

func (pub publisher) Publish() error {
	log.Printf("INFO: publish (%s) message for %s", pub.id, pub.message)
	return pub.broker.Publish(pub.channel, pub.message)
}

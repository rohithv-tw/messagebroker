package Publisher

import (
	"fmt"
	"github.com/google/uuid"
	"message-broker/Log"
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
	Log.Current().LogInfo(
		fmt.Sprintf("Publisher Id (%s), published message : %s", pub.id, pub.message))
	return pub.broker.Publish(pub.channel, pub.message)
}

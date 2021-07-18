package Publisher

import (
	"fmt"
	"github.com/google/uuid"
	Broker2 "message-broker/Broker"
	"message-broker/Log"
	Message2 "message-broker/Message"
)

type publisher struct {
	channel string
	message Message2.IMessage
	broker  Broker2.IBroker
	id      uuid.UUID
}

func Create(broker Broker2.IBroker) IPublisher {
	return publisher{
		broker: broker,
		id:     uuid.New(),
	}
}

func (pub publisher) SetContext(channel string, message Message2.IMessage) IPublisher {
	pub.message = message
	pub.channel = channel
	return pub
}

func (pub publisher) Publish() error {
	Log.Current().LogInfo(
		fmt.Sprintf("Publisher Id (%s), published message : %s", pub.id, pub.message))
	return pub.broker.Publish(pub.channel, pub.message)
}

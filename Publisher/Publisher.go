package Publisher

import (
	"fmt"
	"github.com/google/uuid"
	"message-broker/Broker"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
)

type publisher struct {
	config Config.IConfig
	message Message.IMessage
	broker  Broker.IBroker
	id      uuid.UUID
}

func Create(broker Broker.IBroker, config Config.IConfig) IPublisher {
	return publisher{
		broker: broker,
		id:     uuid.New(),
		config: config,
	}
}

func (pub publisher) SetMessage(message Message.IMessage) IPublisher {
	pub.message = message
	return pub
}

func (pub publisher) Publish() error {
	Log.Current().LogInfo(
		fmt.Sprintf("PublisherId (%s), published message : %s", pub.id, pub.message))
	return pub.broker.Publish(pub.message, pub.config)
}

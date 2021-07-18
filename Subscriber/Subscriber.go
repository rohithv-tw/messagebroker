package Subscriber

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"message-broker/Broker"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
)

type subscriber struct {
	config Config.IConfig
	id      uuid.UUID
	broker  Broker.IBroker
	channel <-chan *bytes.Buffer
}

func Create(broker Broker.IBroker, config Config.IConfig) ISubscriber {
	return subscriber{
		broker: broker,
		id:     uuid.New(),
		config: config,
	}
}

func (s subscriber) Subscribe() error {
	var err error
	s.channel, err = s.broker.Subscribe(s.config)
	if err == nil {
		Log.Current().LogInfo(
			fmt.Sprintf("SubscriberId (%s), subscribe started", s.id))
		for mes := range s.channel {
			message := Message.Create(mes.Bytes())
			Log.Current().LogInfo(
				fmt.Sprintf("SubscriberId (%s), received message : %s\n", s.id, message))
		}
	}
	return err
}

func (s subscriber) Unsubscribe() error {
	return s.broker.Unsubscribe(s.config)
}

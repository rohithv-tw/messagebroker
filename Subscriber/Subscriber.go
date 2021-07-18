package Subscriber

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	Broker2 "message-broker/Broker"
	"message-broker/Log"
	Message2 "message-broker/Message"
)

type subscriber struct {
	id      uuid.UUID
	broker  Broker2.IBroker
	channel <-chan *bytes.Buffer
}

func Create(broker Broker2.IBroker) ISubscriber {
	return subscriber{
		broker: broker,
		id:     uuid.New(),
	}
}

func (s subscriber) Subscribe(channel string) error {
	var err error
	s.channel, err = s.broker.Subscribe(channel)
	if err == nil {
		Log.Current().LogInfo(
			fmt.Sprintf("Subscriber Id (%s), subscribe started", s.id))
		for mes := range s.channel {
			message := Message2.Create(mes.Bytes())
			Log.Current().LogInfo(
				fmt.Sprintf("SubscriberId (%s), received message : %s\n", s.id, message))
		}
	}
	return err
}

func (s subscriber) Unsubscribe(channel string) error {
	return s.broker.Unsubscribe(channel)
}
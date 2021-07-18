package Subscriber

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"log"
	"message-broker/messagebroker/Broker"
)

type subscriber struct {
	id      uuid.UUID
	broker  Broker.IBroker
	channel <-chan *bytes.Buffer
}

func Create(broker Broker.IBroker) ISubscriber {
	return subscriber{
		broker: broker,
		id:     uuid.New(),
	}
}

func (s subscriber) Subscribe(channel string) error {
	var err error
	s.channel, err = s.broker.Subscribe(channel)
	if err == nil {
		log.Printf("INFO: subscribe (%s) message started", s.id)
		for mes := range s.channel {
			fmt.Printf("Subsciber (%s) Received message : %s\n", s.id, mes.String())
		}
	}
	return err
}

func (s subscriber) Unsubscribe(channel string) error {
	return s.broker.Unsubscribe(channel)
}

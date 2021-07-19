package Subscriber

import (
	"fmt"
	"github.com/google/uuid"
	"message-broker/Broker"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
)

type subscriber struct {
	config Config.IConfig
	id     uuid.UUID
	broker Broker.IBroker
}

func Create(broker Broker.IBroker, config Config.IConfig) ISubscriber {
	return subscriber{
		broker: broker,
		id:     uuid.New(),
		config: config,
	}
}

func (s subscriber) Subscribe() error {
	switch s.broker.GetType() {
	case Config.Inmemory:
		return s.SubscribeInMemory()
	case Config.Etcd:
		return s.SubscribeEtcd()
	}
	return fmt.Errorf("unsupported Broker type = %s\n", s.broker.GetType())
}

func (s *subscriber) SubscribeInMemory() error {
	channel, err := s.broker.SubscribeInMemory(s.config)
	if err == nil {
		Log.Current().LogInfo(
			fmt.Sprintf("SubscriberId (%s), subscribe started", s.id))
		for mes := range channel {
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

func (s subscriber) SubscribeEtcd() error {
	channel := s.broker.SubscribeEtcd(s.config)
	Log.Current().LogInfo(
		fmt.Sprintf("SubscriberId (%s), subscribe started", s.id))
	for response := range channel {
		for _, event := range response.Events {
			Log.Current().LogInfo(
				fmt.Sprintf("SubscriberId (%s), %s executed on %q with value %q\n",
					s.id, event.Type, event.Kv.Key, event.Kv.Value))
		}
	}
	return nil
}

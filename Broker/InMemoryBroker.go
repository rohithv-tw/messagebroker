package Broker

import (
	"bytes"
	"fmt"
	"message-broker/Config"
	"message-broker/Message"
	"sync"
)

const defaultBufferValue = 100

type inMemoryBroker struct {
	channels sync.Map // map[string] chan *bytes.Buffer
}

func (i *inMemoryBroker) Subscribe(config Config.IConfig) (<-chan *bytes.Buffer, error) {
	var channelSubscribed chan *bytes.Buffer
	ch, ok := i.channels.Load(config.GetTopic())
	if ok {
		channelSubscribed = ch.(chan *bytes.Buffer)
	} else {
		channelSubscribed = make(chan *bytes.Buffer, defaultBufferValue)
		i.channels.Store(config.GetTopic(), channelSubscribed)
	}
	return channelSubscribed, nil
}

func (i *inMemoryBroker) Unsubscribe(config Config.IConfig) error {
	ch, ok := i.channels.Load(config.GetTopic())
	if ok {
		actualChan := ch.(chan *bytes.Buffer)
		close(actualChan)
		i.channels.Delete(config.GetTopic())
		return nil
	}
	return fmt.Errorf("cannot find channel %s", config.GetTopic())
}

func (i *inMemoryBroker) Publish(message Message.IMessage, config Config.IConfig) error {
	ch, ok := i.channels.Load(config.GetTopic())
	if ok {
		actualChan := ch.(chan *bytes.Buffer)
		actualChan <- message.CreateMessage()
		return nil
	}
	return fmt.Errorf("cannot find channel %s", config.GetTopic())
}

func (i *inMemoryBroker) Close() error {
	i.channels.Range(func(key interface{}, value interface{}) bool {
		ch := value.(chan *bytes.Buffer)
		close(ch)
		i.channels.Delete(key)
		return true
	})
	return nil
}

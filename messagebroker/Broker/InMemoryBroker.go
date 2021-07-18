package Broker

import (
	"bytes"
	"fmt"
	"message-broker/messagebroker/Message"
	"sync"
)

const defaultBufferValue = 100

type inMemoryBroker struct {
	channels sync.Map // map[string] chan *bytes.Buffer
}

func (i *inMemoryBroker) Subscribe(channel string) (chan *bytes.Buffer, error) {
	var channelSubscribed chan *bytes.Buffer
	ch, ok := i.channels.Load(channel)
	if ok {
		channelSubscribed = ch.(chan *bytes.Buffer)
	} else {
		channelSubscribed = make(chan *bytes.Buffer, defaultBufferValue)
		i.channels.Store(channel, channelSubscribed)
	}
	return channelSubscribed, nil
}

func (i *inMemoryBroker) Unsubscribe(channel string) error {
	ch, ok := i.channels.Load(channel)
	if ok {
		actualChan := ch.(chan *bytes.Buffer)
		close(actualChan)
		i.channels.Delete(channel)
		return nil
	}
	return fmt.Errorf("cannot find channel %s", channel)
}

func (i *inMemoryBroker) Publish(channel string, message Message.IMessage) error {
	ch, ok := i.channels.Load(channel)
	if ok {
		actualChan := ch.(chan *bytes.Buffer)
		actualChan <- message.CreateMessage()
		return nil
	}
	return fmt.Errorf("cannot find channel %s", channel)
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

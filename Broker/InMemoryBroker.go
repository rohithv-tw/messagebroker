package Broker

import (
	"bytes"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
	"sync"
)

const defaultBufferValue = 100

type inMemoryBroker struct {
	channels sync.Map // map[string] chan *bytes.Buffer
}

func (i *inMemoryBroker) GetType() string {
	return Config.Inmemory
}

func (i *inMemoryBroker) SubscribeInMemory(config Config.IConfig) (<-chan *bytes.Buffer, error) {
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

func (i *inMemoryBroker) SubscribeEtcd(config Config.IConfig) clientv3.WatchChan {
	Log.Current().LogWarn("Not implemented SubscribeEtcd for InMemoryBroker")
	return nil
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

func (i *inMemoryBroker) Close() error {
	i.channels.Range(func(key interface{}, value interface{}) bool {
		ch := value.(chan *bytes.Buffer)
		close(ch)
		i.channels.Delete(key)
		return true
	})
	return nil
}

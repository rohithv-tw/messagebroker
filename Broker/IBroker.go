package Broker

import (
	"bytes"
	"go.etcd.io/etcd/client/v3"
	"message-broker/Config"
	"message-broker/Message"
)

type IBroker interface {
	GetType() string
	Publish(message Message.IMessage, config Config.IConfig) error
	SubscribeInMemory(config Config.IConfig) (<-chan *bytes.Buffer, error)
	SubscribeEtcd(config Config.IConfig) clientv3.WatchChan
	Unsubscribe(config Config.IConfig) error
	Close() error
}

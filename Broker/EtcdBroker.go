package Broker

import (
	"bytes"
	"message-broker/Config"
	"message-broker/Message"
)

type etcdBroker struct {
}

func (e *etcdBroker) Publish(message Message.IMessage, config Config.IConfig) error {
	panic("implement me")
}

func (e *etcdBroker) Subscribe(Config.IConfig) (<-chan *bytes.Buffer, error) {
	panic("implement me")
}

func (e *etcdBroker) Unsubscribe(Config.IConfig) error {
	panic("implement me")
}

func (e *etcdBroker) Close() error {
	panic("implement me")
}

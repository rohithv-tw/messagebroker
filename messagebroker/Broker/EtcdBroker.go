package Broker

import (
	"bytes"
	"message-broker/messagebroker/Message"
)

type etcdBroker struct {
}

func (etcd *etcdBroker) Subscribe(channel string) (chan *bytes.Buffer, error) {
	panic("Implement me")
}

func (etcd *etcdBroker) Publish(channel string, message Message.IMessage) error {
	panic("implement me")
}

func (etcd *etcdBroker) Unsubscribe(channel string) error {
	panic("implement me")
}

func (etcd *etcdBroker) Close() error {
	panic("implement me")
}

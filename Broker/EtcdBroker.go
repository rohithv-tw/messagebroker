package Broker

import (
	"bytes"
	Message2 "message-broker/Message"
)

type etcdBroker struct {
}

func (etcd *etcdBroker) Subscribe(channel string) (<-chan *bytes.Buffer, error) {
	panic("Implement me")
}

func (etcd *etcdBroker) Publish(channel string, message Message2.IMessage) error {
	panic("implement me")
}

func (etcd *etcdBroker) Unsubscribe(channel string) error {
	panic("implement me")
}

func (etcd *etcdBroker) Close() error {
	panic("implement me")
}

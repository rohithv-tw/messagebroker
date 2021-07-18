package Broker

import (
	"bytes"
	Message2 "message-broker/Message"
)

type IBroker interface {
	Publish(channel string, message Message2.IMessage) error
	Subscribe(channel string) (<-chan *bytes.Buffer, error)
	Unsubscribe(channel string) error
	Close() error
}

package Broker

import (
	"bytes"
	"message-broker/messagebroker/Message"
)

type IBroker interface {
	Publish(channel string, message Message.IMessage) error
	Subscribe(channel string) (chan *bytes.Buffer, error)
	Unsubscribe(channel string) error
	Close() error
}

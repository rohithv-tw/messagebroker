package Broker

import (
	"bytes"
	"message-broker/Config"
	"message-broker/Message"
)

type IBroker interface {
	Publish(message Message.IMessage, config Config.IConfig) error
	Subscribe(config Config.IConfig) (<-chan *bytes.Buffer, error)
	Unsubscribe(config Config.IConfig) error
	Close() error
}

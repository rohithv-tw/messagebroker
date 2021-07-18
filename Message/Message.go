package Message

import (
	"bytes"
)

type messageWrapper struct {
	messageBytes []byte
}

func Create(message []byte) IMessage {
	return messageWrapper{
		messageBytes: message,
	}
}

func (m messageWrapper) CreateMessage() *bytes.Buffer {
	return bytes.NewBuffer(m.messageBytes)
}

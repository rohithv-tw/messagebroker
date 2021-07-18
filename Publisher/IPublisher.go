package Publisher

import (
	Message2 "message-broker/Message"
)

type IPublisher interface {
	SetContext(channel string, message Message2.IMessage) IPublisher
	Publish() error
}

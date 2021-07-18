package Publisher

import (
	"message-broker/Message"
)

type IPublisher interface {
	SetContext(channel string, message Message.IMessage) IPublisher
	Publish() error
}

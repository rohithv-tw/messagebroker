package Publisher

import (
	"message-broker/Message"
)

type IPublisher interface {
	SetMessage(message Message.IMessage) IPublisher
	Publish() error
}

package Publisher

import "message-broker/messagebroker/Message"

type IPublisher interface {
	SetContext(channel string, message Message.IMessage) IPublisher
	Publish() error
}

package Message

import "bytes"

type IMessage interface {
	CreateMessage() *bytes.Buffer
}

package Subscriber

type ISubscriber interface {
	Subscribe(channel string) error
	Unsubscribe(channel string) error
}

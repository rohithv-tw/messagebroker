package Subscriber

type ISubscriber interface {
	Subscribe() error
	Unsubscribe() error
}

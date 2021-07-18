package main

import (
	"fmt"
	"log"
	"message-broker/messagebroker/Broker"
	"message-broker/messagebroker/Message"
	"message-broker/messagebroker/Publisher"
	"message-broker/messagebroker/Subscriber"
	"sync/atomic"
)

var increment int32

const (
	count       = 1
	channelName = "channel"
)

func main() {
	brok, err := Broker.Create(Broker.InMemory)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt32(&increment, 1)
			message := fmt.Sprintf("Increment - %d", atomic.LoadInt32(&increment))
			messageWrapper := Message.Create([]byte(message))
			pub := Publisher.Create(brok)
			pub = pub.SetContext(channelName, messageWrapper)
			_ = pub.Publish()
		}
	}()

	var subs []Subscriber.ISubscriber
	for i := 0; i < count; i++ {
		subs = append(subs, Subscriber.Create(brok))
		subs[i].Subscribe(channelName)
	}

	defer func() {
		for _, sub := range subs {
			sub.Unsubscribe(channelName)
		}
		_ = brok.Close()
	}()

}

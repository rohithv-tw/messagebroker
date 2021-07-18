package main

import (
	"fmt"
	"message-broker/Log"
	"message-broker/messagebroker/Broker"
	"message-broker/messagebroker/Message"
	"message-broker/messagebroker/Publisher"
	"message-broker/messagebroker/Subscriber"
	"time"
)

const (
	count       = 1
	channelName = "channel"
)

func main() {
	brok, err := Broker.Create(Broker.InMemory)
	logError(err)

	go func() {
		for i := 0; i < count; i++ {
			message := fmt.Sprintf(
				"Generated Message -> %s at %d", time.Now().String(), time.Now().UnixNano())
			messageWrapper := Message.Create([]byte(message))
			pub := Publisher.Create(brok)
			pub = pub.SetContext(channelName, messageWrapper)
			err = pub.Publish()
			logError(err)
		}
	}()

	var subs []Subscriber.ISubscriber
	for i := 0; i < count; i++ {
		subs = append(subs, Subscriber.Create(brok))
		err = subs[i].Subscribe(channelName)
		logError(err)
	}

	defer func() {
		for _, sub := range subs {
			err = sub.Unsubscribe(channelName)
			logError(err)
		}
		err = brok.Close()
		logError(err)
	}()
}

func logError(err error){
	if err!=nil{
		Log.Current().LogError(err)
	}
}

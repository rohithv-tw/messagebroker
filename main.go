package main

import (
	"flag"
	"fmt"
	"message-broker/Broker"
	"message-broker/Log"
	"message-broker/Message"
	"message-broker/Publisher"
	"message-broker/Subscriber"
	"time"
)

const (
	count     = 1
	topicName = "channel"
)

func main() {
	brokerType := flag.String("brokerType", Broker.Inmemory, "Broker Type")
	flag.Parse()

	brok, err := Broker.Create(*brokerType)
	logError(err)

	go func() {
		for i := 0; i < count; i++ {
			message := fmt.Sprintf(
				"Generated Message -> %s at %d", time.Now().String(), time.Now().UnixNano())
			messageWrapper := Message.Create([]byte(message))
			pub := Publisher.Create(brok)
			pub = pub.SetContext(topicName, messageWrapper)
			err = pub.Publish()
			logError(err)
		}
	}()

	var subs []Subscriber.ISubscriber
	for i := 0; i < count; i++ {
		subs = append(subs, Subscriber.Create(brok))
		err = subs[i].Subscribe(topicName)
		logError(err)
	}

	defer func() {
		for _, sub := range subs {
			err = sub.Unsubscribe(topicName)
			logError(err)
		}
		err = brok.Close()
		logError(err)
	}()
}

func logError(err error) {
	if err != nil {
		Log.Current().LogError(err)
	}
}

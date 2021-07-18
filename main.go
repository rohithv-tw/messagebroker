package main

import (
	"flag"
	"fmt"
	"message-broker/Broker"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
	"message-broker/Publisher"
	"message-broker/Subscriber"
	"time"
)

const count = 1

func main() {
	brokerType := flag.String("brokerType", Broker.Inmemory, "Broker Type")
	flag.Parse()

	config, err := Config.Create(*brokerType)
	logErrorWhenApplicable(err)

	brok, err := Broker.Create(*brokerType)
	logErrorWhenApplicable(err)

	go func() {
		for i := 0; i < count; i++ {
			message := fmt.Sprintf(
				"Generated Message -> %s at %d", time.Now().String(), time.Now().UnixNano())
			messageWrapper := Message.Create([]byte(message))
			pub := Publisher.Create(brok, config)
			pub = pub.SetMessage(messageWrapper)
			err = pub.Publish()
			logErrorWhenApplicable(err)
		}
	}()

	var subs []Subscriber.ISubscriber
	for i := 0; i < count; i++ {
		subs = append(subs, Subscriber.Create(brok, config))
		err = subs[i].Subscribe()
		logErrorWhenApplicable(err)
	}

	defer func() {
		for _, sub := range subs {
			err = sub.Unsubscribe()
			logErrorWhenApplicable(err)
		}
		err = brok.Close()
		logErrorWhenApplicable(err)
	}()
}

func logErrorWhenApplicable(err error) {
	if err != nil {
		Log.Current().LogError(err)
	}
}

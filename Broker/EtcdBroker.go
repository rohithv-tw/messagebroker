package Broker

import (
	"bytes"
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
	"sync"
	"time"
)

type etcdBroker struct {
	client *clientv3.Client
}

func (e *etcdBroker) GetType() string {
	return Config.Etcd
}

func (e *etcdBroker) SubscribeInMemory(Config.IConfig) (<-chan *bytes.Buffer, error) {
	Log.Current().LogWarn("Not implemented SubscribeInMemory for EtcdBroker")
	return nil, nil
}

func (e *etcdBroker) SubscribeEtcd(config Config.IConfig) clientv3.WatchChan {
	e.initialiseClient(config, "subscribe()")
	return e.client.Watch(context.Background(), config.GetTopic())
}

var mutex sync.Mutex

func (e *etcdBroker) initialiseClient(config Config.IConfig, event string) {
	mutex.Lock()
	defer mutex.Unlock()
	var err error
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + config.GetHost()},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(fmt.Sprintf("connected to etcd %s on %s", e.client.Endpoints(), event))
	}
}

func (e *etcdBroker) Publish(message Message.IMessage, config Config.IConfig) error {
	e.initialiseClient(config, "publish()")
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(e.client)
	res, err := e.client.Put(context.Background(), config.GetTopic(), message.CreateMessage().String())
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(res.PrevKv.String())
	}
	return nil
}

func (e *etcdBroker) Unsubscribe(config Config.IConfig) error {
	e.initialiseClient(config, "unsubscribe()")
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(e.client)
	return nil
}

func (e *etcdBroker) Close() error {
	func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(e.client)
	return nil
}

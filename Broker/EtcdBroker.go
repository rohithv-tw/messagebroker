package Broker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"message-broker/Config"
	"message-broker/Log"
	"message-broker/Message"
	"time"
)

type etcdBroker struct {
}

func (e *etcdBroker) GetType() string {
	return Config.Etcd
}

func (e *etcdBroker) SubscribeInMemory(Config.IConfig) (<-chan *bytes.Buffer, error) {
	Log.Current().LogWarn("Not implemented SubscribeInMemory for EtcdBroker")
	return nil, nil
}

func (e *etcdBroker) SubscribeEtcd(config Config.IConfig) clientv3.WatchChan {
	client, err := e.initialiseClient(config.GetHost())
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(client)
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(
			fmt.Sprintf("connected to etcd %s on %s", client.Endpoints(), "subscribe()"))
	}
	return client.Watch(context.Background(), config.GetTopic())
}

func (e *etcdBroker) initialiseClient(host string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://" + host},
		DialTimeout: 5 * time.Second,
	})
}

func (e *etcdBroker) Publish(message Message.IMessage, config Config.IConfig) error {
	client, err := e.initialiseClient(config.GetHost())
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(client)
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(
			fmt.Sprintf("connected to etcd %s on %s", client.Endpoints(), "publish()"))
	}
	var res *clientv3.PutResponse
	res, err = client.Put(context.Background(), config.GetTopic(), message.CreateMessage().String())
	if err != nil {
		Log.Current().LogError(err)
	} else {
		serializedData, _ := json.Marshal(res)
		Log.Current().LogInfo(string(serializedData))
	}
	return nil
}

func (e *etcdBroker) Unsubscribe(config Config.IConfig) error {
	client, err := e.initialiseClient(config.GetHost())
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(client)
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(
			fmt.Sprintf("connected to etcd %s on %s", client.Endpoints(), "Unsubscribe()"))
	}
	return nil
}

func (e *etcdBroker) Close(config Config.IConfig) error {
	client, err := e.initialiseClient(config.GetHost())
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			Log.Current().LogError(err)
		}
	}(client)
	if err != nil {
		Log.Current().LogError(err)
	} else {
		Log.Current().LogInfo(
			fmt.Sprintf("connected to etcd %s on %s", client.Endpoints(), "close()"))
	}
	return nil
}

package Config

import (
	"fmt"
	"message-broker/Broker"
)

type config struct {
	host  string
	topic string
}

const topic = "channel"

func Create(brokerType string) (IConfig, error) {
	switch brokerType {
	case Broker.Inmemory:
		return &config{host: "", topic: topic}, nil
	case Broker.Etcd:
		return &config{host: "192.168.99.100:2379", topic: topic}, nil
	default:
		return nil, fmt.Errorf("unsupported Broker type = %d", brokerType)
	}
}

func (c *config) GetHost() string {
	return c.host
}

func (c *config) GetTopic() string {
	return c.topic
}

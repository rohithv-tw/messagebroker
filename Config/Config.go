package Config

import (
	"fmt"
	"strings"
)

type config struct {
	host  string
	topic string
}

const topic = "channel"

func Create(brokerType string) (IConfig, error) {
	brokerType = strings.ToLower(brokerType)
	switch brokerType {
	case Inmemory:
		return &config{host: "", topic: topic}, nil
	case Etcd:
		return &config{host: "127.0.0.1:2379", topic: topic}, nil
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

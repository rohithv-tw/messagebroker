package Broker

import (
	"fmt"
	"message-broker/Config"
	"strings"
	"sync"
)

func Create(brokerType string) (IBroker, error) {
	brokerType = strings.ToLower(brokerType)
	switch brokerType {
	case Config.Inmemory:
		return &inMemoryBroker{
			channels: sync.Map{},
		}, nil
	case Config.Etcd:
		return &etcdBroker{}, nil
	default:
		return nil, fmt.Errorf("unsupported Broker type = %d", brokerType)
	}
}

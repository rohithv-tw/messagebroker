package Broker

import (
	"fmt"
	"strings"
	"sync"
)

func Create(brokerType string) (IBroker, error) {
	brokerType = strings.ToLower(brokerType)
	switch brokerType {
	case Inmemory:
		return &inMemoryBroker{
			channels: sync.Map{},
		}, nil
	case Etcd:
		return &etcdBroker{}, nil
	default:
		return nil, fmt.Errorf("unsupported Broker type = %d", brokerType)
	}
}

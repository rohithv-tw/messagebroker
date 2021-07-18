package Broker

import (
	"fmt"
	"sync"
)

var brokerMap sync.Map

func Create(brokerType int) (IBroker, error) {
	result, _ := brokerMap.LoadOrStore(brokerType, func() IBroker {
		res, _ := createInstance(brokerType)
		return res
	}())
	return result.(IBroker), nil
}

func createInstance(brokerType int) (IBroker, error) {
	switch brokerType {
	case InMemory:
		return &inMemoryBroker{
			channels: sync.Map{},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported Broker type = %d", brokerType)
	}
}

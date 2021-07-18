package Broker

import (
	"fmt"
	"sync"
)

var instance IBroker
var once sync.Once

func Create(brokerType int) (IBroker, error) {
	var err error = nil
	once.Do(func() {
		switch brokerType {
		case InMemory:
			instance = &inMemoryBroker{
				channels: sync.Map{},
			}
		case Etcd:{
			instance = &etcdBroker{}
		}
		default:
			err = fmt.Errorf("unsupported Broker type = %d", brokerType)
		}
	})
	return instance, err
}

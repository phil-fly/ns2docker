package ns2docker

import (
	"github.com/docker/docker/api/types"
	"sync"
)

type NsCache struct {
	Ns2Container map[string]types.Container
	PoolMutex	*sync.RWMutex
}

var DockerNsCache *NsCache

func init() {
	DockerNsCache = NewNsCache()
}

func NewNsCache() (ns *NsCache) {
	return &NsCache{
		Ns2Container:make(map[string]types.Container),
		PoolMutex:new(sync.RWMutex),
	}
}

func (nc *NsCache)Put(key string,value types.Container) {
	nc.PoolMutex.Lock()
	nc.Ns2Container[key] = value
	nc.PoolMutex.Unlock()
}

func (nc *NsCache)Get(key string) (types.Container,bool) {
	nc.PoolMutex.RLock()
	defer nc.PoolMutex.RUnlock()
	value,ok := nc.Ns2Container[key]
	return value,ok
}

func (nc *NsCache)Del(key string) {
	nc.PoolMutex.Lock()
	defer nc.PoolMutex.Unlock()
	delete(nc.Ns2Container, key)
}

func (nc *NsCache)Clear() {
	for k :=range nc.Ns2Container {
		delete(nc.Ns2Container,k)
	}
}
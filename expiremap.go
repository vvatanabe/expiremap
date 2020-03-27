package expiremap

import (
	"sync"
	"time"
)

type Map struct {
	syncMap       sync.Map
	defaultExpire time.Duration
}

func (m *Map) SetDefaultExpire(expire time.Duration) {
	m.defaultExpire = expire
}

func (m *Map) Delete(key interface{}) {
	m.syncMap.Delete(key)
}

func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	return m.syncMap.Load(key)
}

func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	return m.LoadOrStoreWithExpire(key, value, m.defaultExpire)
}

func (m *Map) LoadOrStoreWithExpire(key, value interface{}, expire time.Duration) (actual interface{}, loaded bool) {
	actual, loaded = m.syncMap.LoadOrStore(key, value)
	if !loaded {
		m.setExpire(key, expire)
	}
	return
}

func (m *Map) Range(f func(key, value interface{}) bool) {
	m.syncMap.Range(f)
}

func (m *Map) Store(key, value interface{}) {
	m.StoreWithExpire(key, value, m.defaultExpire)
}

func (m *Map) StoreWithExpire(key, value interface{}, expire time.Duration) {
	m.syncMap.Store(key, value)
	m.setExpire(key, expire)
}

func (m *Map) setExpire(key interface{}, expire time.Duration) {
	if expire == 0 {
		return
	}
	go func() {
		<-time.Tick(expire)
		m.syncMap.Delete(key)
	}()
}

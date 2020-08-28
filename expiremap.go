package expiremap

import (
	"sync"
	"time"
)

// Map is synchronization map with expiration date (extended sync.Map).
type Map struct {
	syncMap       sync.Map
	defaultExpire time.Duration
}

// SetDefaultExpire sets default expiration for value in Map.
func (m *Map) SetDefaultExpire(expire time.Duration) {
	m.defaultExpire = expire
}

// Delete deletes the value for a key.
func (m *Map) Delete(key interface{}) {
	m.syncMap.Delete(key)
}

// Load returns the value in expiration stored in the map for a key,
// or nil if no value is present.
func (m *Map) Load(key interface{}) (value interface{}, ok bool) {
	return m.syncMap.Load(key)
}

// LoadOrStore returns the existing value in expiration for the key
// if present.
// Otherwise, it stores with expiration and returns the given value.
func (m *Map) LoadOrStore(key, value interface{}, opts ...Option) (actual interface{}, loaded bool) {
	actual, loaded = m.syncMap.LoadOrStore(key, value)
	if !loaded {
		opt := &Options{
			Expire:      m.defaultExpire,
			ExpiredFunc: func() {},
		}
		for _, optFunc := range opts {
			optFunc(opt)
		}
		m.setExpire(key, opt.Expire, opt.ExpiredFunc)
	}
	return actual, loaded
}

// Range calls f sequentially for each key and value in expiration
// present in the map.
func (m *Map) Range(f func(key, value interface{}) bool) {
	m.syncMap.Range(f)
}

// Store sets the value for a key with default expiration and some options.
func (m *Map) Store(key, value interface{}, opts ...Option) {
	opt := &Options{
		Expire:      m.defaultExpire,
		ExpiredFunc: func() {},
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	m.syncMap.Store(key, value)
	m.setExpire(key, opt.Expire, opt.ExpiredFunc)
}

func (m *Map) setExpire(key interface{}, expire time.Duration, f func()) {
	if expire == 0 {
		return
	}
	go func() {
		<-time.Tick(expire)
		m.syncMap.Delete(key)
		f()
	}()
}

type Options struct {
	Expire      time.Duration
	ExpiredFunc func()
}

type Option func(*Options)

func Expire(expire time.Duration) Option {
	return func(args *Options) {
		args.Expire = expire
	}
}

func ExpiredFunc(f func()) Option {
	return func(args *Options) {
		args.ExpiredFunc = f
	}
}

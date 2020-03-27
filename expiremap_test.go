package expiremap

import (
	"testing"
	"time"
)

func TestMap_Store_And_Load_And_Expired(t *testing.T) {
	var m Map

	wantKey := "foo"
	wantValue := "bar"

	m.StoreWithExpire(wantKey, wantValue, time.Second/2)
	v, ok := m.Load(wantKey)
	if !ok {
		t.Errorf("Map.Load() not exixts key=%s", wantKey)
		return
	}
	if got := v.(string); got != wantValue {
		t.Errorf("Map.Load() gotValue = %v, want %v", got, wantValue)
		return
	}

	time.Sleep(time.Second)

	_, ok = m.Load(wantKey)
	if ok {
		t.Errorf("Map.Load() exixts key=%s", wantKey)
		return
	}
}

func TestMap_Store_And_Load_And_Expired_Default(t *testing.T) {
	var m Map
	m.SetDefaultExpire(time.Second / 2)

	wantKey := "foo"
	wantValue := "bar"

	m.Store(wantKey, wantValue)
	v, ok := m.Load(wantKey)
	if !ok {
		t.Errorf("Map.Load() not exixts key=%s", wantKey)
		return
	}
	if got := v.(string); got != wantValue {
		t.Errorf("Map.Load() gotValue = %v, want %v", got, wantValue)
		return
	}

	time.Sleep(time.Second)

	_, ok = m.Load(wantKey)
	if ok {
		t.Errorf("Map.Load() exixts key=%s", wantKey)
		return
	}
}

func TestMap_LoadOrStore_And_Load_And_Expired(t *testing.T) {
	var m Map

	wantKey := "foo"
	wantValue := "bar"

	m.LoadOrStoreWithExpire(wantKey, wantValue, time.Second/2)
	v, ok := m.Load(wantKey)
	if !ok {
		t.Errorf("Map.Load() not exixts key=%s", wantKey)
		return
	}
	if got := v.(string); got != wantValue {
		t.Errorf("Map.Load() gotValue = %v, want %v", got, wantValue)
		return
	}

	time.Sleep(time.Second)

	_, ok = m.Load(wantKey)
	if ok {
		t.Errorf("Map.Load() exixts key=%s", wantKey)
		return
	}
}

func TestMap_LoadOrStore_And_Expired_Default(t *testing.T) {
	var m Map
	m.SetDefaultExpire(time.Second / 2)

	wantKey := "foo"
	wantValue := "bar"

	m.LoadOrStore(wantKey, wantValue)
	v, ok := m.Load(wantKey)
	if !ok {
		t.Errorf("Map.Load() not exixts key=%s", wantKey)
		return
	}
	if got := v.(string); got != wantValue {
		t.Errorf("Map.Load() gotValue = %v, want %v", got, wantValue)
		return
	}

	time.Sleep(time.Second)

	_, ok = m.Load(wantKey)
	if ok {
		t.Errorf("Map.Load() exixts key=%s", wantKey)
		return
	}
}

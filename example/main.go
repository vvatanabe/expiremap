package main

import (
	"fmt"
	"time"

	"github.com/vvatanabe/expiremap"
)

func main() {
	var m expiremap.Map

	// SetDefaultExpire sets default expiration for value in Map.
	m.SetDefaultExpire(time.Second / 2)

	// Store sets the value for a key with default expiration.
	m.Store("key1", "foo")

	// Load returns the value in expiration stored in the map for a key,
	// or nil if no value is present.
	v, ok := m.Load("key1")
	if !ok {
		return
	}
	fmt.Println("key1:", v)

	// Store with expire sets the value for a key with expiration.
	m.Store("key2", "bar", expiremap.Expire(time.Second/2))
	v, ok = m.Load("key2")
	if !ok {
		return
	}
	fmt.Println("key2:", v)

	// Wait for expiration
	time.Sleep(time.Second)

	_, ok = m.Load("key1")
	if !ok {
		fmt.Println("key1 expired")
	}
	_, ok = m.Load("key2")
	if !ok {
		fmt.Println("key2 expired")
	}
}

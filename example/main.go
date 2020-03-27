package main

import (
	"fmt"
	"time"

	"github.com/vvatanabe/expiremap"
)

func main() {
	var m expiremap.Map
	m.SetDefaultExpire(time.Second / 2)

	m.Store("key1", "foo")
	v, ok := m.Load("key1")
	if !ok {
		return
	}
	fmt.Println("key1:", v)

	m.StoreWithExpire("key2", "bar", time.Second/2)
	v, ok = m.Load("key2")
	if !ok {
		return
	}
	fmt.Println("key2:", v)

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

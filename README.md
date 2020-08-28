# expiremap ![](https://github.com/vvatanabe/expiremap/workflows/Go/badge.svg)

synchronization map with expiration date (extended sync.Map)

## Description

`expiremap.Map` provides a func that is compatible with `sync.Map` and an extended func.

## Installation

This package can be installed with the go get command:

```
$ go get github.com/vvatanabe/expiremap
```

## Usage

```go
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
	m.Store("key2", "bar", expiremap.Expire(time.Second / 2))
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
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

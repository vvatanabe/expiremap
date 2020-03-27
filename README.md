# expiremap ![](https://github.com/vvatanabe/expiremap/workflows/Go/badge.svg)

synchronization map with expiration date (extended sync.Map)

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
```

## Bugs and Feedback

For bugs, questions and discussions please use the Github Issues.

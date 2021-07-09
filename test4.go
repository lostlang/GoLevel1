package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	i int
}

func newCounter() *Counter {
	return &Counter{
		i: 0,
	}
}

func (c *Counter) Count() int {
	c.Lock()
	defer c.Unlock()

	c.i += 1
	return c.i
}

func main() {
	store := newCounter()

	for i := 0; i < 15; i++ {
		go func() {
			fmt.Println(store.Count())
		}()
	}

	fmt.Scanln()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()   // lock to use
	c.c[key]++    // critical section (in one time can use only one goroutine.
	c.mu.Unlock() // unlock to use
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()         // lock to use.
	defer c.mu.Unlock() // unlock to use (will run after critical section).
	return c.c[key]     // critical section (in one time can use only one goroutine.
}

// Multithreading. Synchronization primitives
func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}

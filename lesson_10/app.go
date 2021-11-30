package main

import (
	"fmt"
	"sync"
	"time"
)

// Counter simple mutex
type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

// Inc increment for simple mutex
func (c *Counter) Inc(key string) {
	c.mu.Lock()   // lock to use
	c.c[key]++    // critical section (in one time can use only one goroutine.
	c.mu.Unlock() // unlock to use
	//fmt.Println("done")
}

// Value getter for simple mutex
func (c *Counter) Value(key string) int {
	c.mu.Lock()         // lock to use.
	defer c.mu.Unlock() // unlock to use (will run after critical section).
	return c.c[key]     // critical section (in one time can use only one goroutine.
}

// simpleMutexPractice practice
func simpleMutexPractice() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}

// Counter2 RWMutex practice
type Counter2 struct {
	mu sync.RWMutex
	c  map[string]int
}

// CountMe increment for RWMutex
func (c *Counter2) CountMe() map[string]int {
	c.mu.Lock()         // rw lock
	defer c.mu.Unlock() // rw unlock
	return c.c          // critical section (in one time can use only one goroutine.
}

// CountMeAgain increment for RWMutex
func (c *Counter2) CountMeAgain() map[string]int {
	c.mu.RLock()         // only read lock
	defer c.mu.RUnlock() // only read unlock
	return c.c
}

// Multithreading. Synchronization primitives
func main() {
	simpleMutexPractice()
}

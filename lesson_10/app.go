package main

import (
	"fmt"
	"sync"
	"sync/atomic"
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

// waitGroupSync sync.WaitGroup allow waiting end all goroutines in one moment (together).
func waitGroupSync() {
	var wg sync.WaitGroup

	//wg.Add(1) // will be panic
	for i := 0; i < 10; i++ {
		// will be panic
		/*if (i == 1) {
			wg.Add(1)
		}*/

		wg.Add(1) // can't be less than 0.
		go func(i int) {
			defer wg.Done() // remove from wg.Add -1

			fmt.Printf("%d goroutine working...\n", i)
			time.Sleep(1 * time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("all done")
}

// atomicForThreadSafe atomic counter.
func atomicForThreadSafe() {
	var wg sync.WaitGroup
	var counter uint64
	//var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				//counter++ // lost part numbers. Not thread safe.

				// Always result 10000.
				//mu.Lock()
				//counter++
				//mu.Unlock()

				atomic.AddUint64(&counter, 1) // Always result 10000. Thread safe
			}
		}()
	}

	wg.Wait()
	fmt.Printf("all done. couner: %d\n", counter)
}

// Multithreading. Synchronization primitives
func main() {
	//simpleMutexPractice()

	//waitGroupSync()

	atomicForThreadSafe()
}

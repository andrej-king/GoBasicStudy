package main

import (
	"fmt"
	"time"
)

// Multithreading. Goroutines and channels
func main() {
	testGoroutines()
}

func sayTest(word string) {
	time.Sleep(1 * time.Second) // sleep 1 second
	fmt.Println(word)
}

// testGoroutines Multithreading
func testGoroutines() {
	go sayTest("Some text") // Run in separate thread
	fmt.Println("1")        // Main thread of execution
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(2 * time.Second)
}

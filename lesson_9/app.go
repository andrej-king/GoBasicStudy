package main

import (
	"fmt"
	"time"
)

// Multithreading. Goroutines and channels
func main() {
	//testGoroutines()
	testGoroutines2()
}

func sayText(word string) {
	time.Sleep(1 * time.Second) // sleep 1 second
	fmt.Println(word)
}

// testGoroutines Multithreading
func testGoroutines() {
	go sayText("Hello world") // Run in separate thread
	fmt.Println("1")          // Main thread of execution
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	time.Sleep(2 * time.Second)
}

func sayText2(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		sayText("Some text")
		exit <- i // write in channel
	}
	close(exit) // close channel
}

func testGoroutines2() {
	ch := make(chan int) // created channel
	go sayText2(ch)
	//resultFromChannel := <-ch // get from channel
	for i := range ch { // iterate in channel
		fmt.Println(i)
	}
}

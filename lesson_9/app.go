package main

import (
	"fmt"
	"time"
)

// Multithreading. Goroutines and channels
func main() {
	//testGoroutines()
	//testGoroutines2()
	data := make(chan int)
	exit := make(chan int)

	// anonymous function will auto call.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data) // read data from channel
		}
		exit <- 0 // write data in channel
	}()
	selectOne(data, exit)
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

// selectOne allow read multiple channels
func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x: // if read data from channel
			x += 1
		case <-exit: // if write data in channel
			fmt.Println("exit")
			return
		default:
			fmt.Println("Waiting...")
			time.Sleep(50 * time.Millisecond)
			//return
		}
	}
}

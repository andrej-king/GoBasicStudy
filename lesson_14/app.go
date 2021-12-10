package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// work with context
func main() {
	//testContextWithCancel()
	//testContextWithDeadline()
	//testContextWithTimeout()
	testContextWithValue()
}

// testContextWithValue put value in context
func testContextWithValue() {
	duration := 1500 * time.Millisecond
	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "hello")
	ctx, cancel := context.WithTimeout(ctx, duration)

	//cancel context
	defer cancel()

	doRequest(ctx, "https://google.com")
}

// testContextWithTimeout run context with timeout
func testContextWithTimeout() {
	duration := 1500 * time.Millisecond
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, duration)

	//cancel context
	defer cancel()

	doRequest(ctx, "https://google.com")
}

// testContextWithDeadline run context with time deadline
func testContextWithDeadline() {
	//duration := 500 * time.Millisecond
	duration := 1 * time.Second
	ctx := context.Background()
	d := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(ctx, d)

	//cancel context
	defer cancel()

	doRequest(ctx, "https://google.com")
}

// testContextWithCancel run context with cancel ability (call function "cancel" or when parent close "Done" channel)
func testContextWithCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//defer cancel()

	//cancel context
	go func() {
		err := cancelRequest(ctx)
		if err != nil {
			cancel()
		}
	}()

	doRequest(ctx, "https://google.com")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Response completed, status code = %d", res.StatusCode)
		if ctx.Value("test") != nil {
			fmt.Println()
			fmt.Println("value by 'test' key:", ctx.Value("test"))
		}
	case <-ctx.Done():
		fmt.Println("Request too long.")

	}
}

func contextDescription() {
	// empty contexts
	//context.Background()
	//context.TODO()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancel() // close context or when parent close "Done" channel
}

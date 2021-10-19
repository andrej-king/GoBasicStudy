package main

import (
	"fmt"
)

// Errors
func main2() {
	err := method1()
	if err != nil {
		fmt.Println(err.Unwrap()) // internal error
		fmt.Println(err.Error())  // something goes wrong. error
		//log.Fatal(err.Error())
		return
	}
	fmt.Println("success")
}

// standard go error interface
//type error interface {
//	Error() string
//}

// appError custom app error
type appError struct {
	Err    error
	Custom string
	Field  int
}

func (e *appError) Error() string {
	// custom error
	err := fmt.Errorf("%s error", e.Custom)
	return err.Error()

	// default error
	//return e.Err.Error()
}

// Unwrap error
func (e *appError) Unwrap() error {
	return e.Err
}

// method1 return custom error
func method1() *appError {
	return method2()
}

// method2 return custom error
func method2() *appError {
	return method3() // if error - this error will be move up to next method
}

// method3 return custom error
func method3() *appError {
	// implement business logic
	return &appError{
		Err:    fmt.Errorf("internal error"),
		Custom: "something goes wrong.",
		Field:  56,
	}
}

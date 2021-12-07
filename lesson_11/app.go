package main

import (
	"errors"
	"fmt"
	"log"
)

// nonExistentIndex get panic
func nonExistentIndex() {
	a := []int{1, 2, 3}
	c := a[3] // get panic
	fmt.Println(c)
}

type name struct {
	A, B int
}

func (n *name) method() {
	fmt.Println("ok")
	fmt.Println(n.A)
}

// getPanicByCallMethod
func getPanicByCallMethod() {
	n := &name{1, 2}
	n = nil
	n.method()
}

// callOwnPanic own panic
func callOwnPanic() {
	// defer call before panic
	defer func() {
		fmt.Println("ok")
	}()
	panic("Something went wrong.")
}

// divide check if default panic happened.
func divide(a, b int) {
	// recovery panic
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic happened", err)
		}
	}()

	fmt.Println(a / b)
}

type AppError struct {
	Message string
	Err     error
}

func (ae *AppError) Error() string {
	return ae.Message
}

// divideWithCustomError
func divideWithCustomError(a, b int) {
	// recovery panic
	defer func() {
		var appErr *AppError
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				if errors.As(err.(error), &appErr) {
					fmt.Println("This is App Error handle panic:", err)
				} else {
					fmt.Println("This is Other Error handle panic:", err)
				}
			default:
				// golang panic
				fmt.Println("This is default go panic:", err)
				//panic("Some panic!")
			}
			log.Println("Panic happened:", err)
		}
	}()

	fmt.Println(div(a, b))
}

func div(a, b int) int {
	if b == 0 {
		// first panic type
		//panic(fmt.Errorf("Some error")) // This is Other Error handle panic

		// second panic type
		panic(&AppError{ // This is App Error handle panic
			Message: "This is divide by zero custom error!",
			Err:     nil,
		})

		// third panic type
		//panic("aaa") // This is default go panic
	}
	return a / b
}

// work with 'panics'
func main() {
	//nonExistentIndex()

	//getPanicByCallMethod()

	//callOwnPanic()

	//divide(4, 2) // ok
	//divide(4, 0)                    // call panic "2021/12/01 21:47:09 panic happened runtime error: integer divide by zero"
	//fmt.Println("Code after panic") // get result "Code after panic"

	// call custom panic
	divideWithCustomError(4, 0)
	fmt.Println("Code after panic") // get result "Code after panic"
}

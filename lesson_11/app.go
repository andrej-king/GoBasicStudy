package main

import (
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

// divide check if panic happened.
func divide(a, b int) {
	// recovery panic
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic happened", err)
		}
	}()

	fmt.Println(a / b)
}

// work with 'panics'
func main() {
	//nonExistentIndex()

	//getPanicByCallMethod()

	//callOwnPanic()

	//divide(4, 2) // ok
	divide(4, 0) // call panic
}

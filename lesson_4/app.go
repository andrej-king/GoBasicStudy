package main

import "fmt"

// Functions and methods
func main() {
	//callbackMathExamples()

	// closures
	//orderPrice := totalPrice(1)
	//fmt.Println(orderPrice(1)) // 2 - each time will run inner method
	//fmt.Println(orderPrice(1)) // 3

	// work structures
	p := Point{1, 2}
	fmt.Println(p)                  // {1 2}
	fmt.Println(movePoint(p, 1, 1)) // {2 3}
	fmt.Println(p)                  // {1 2} - because "movePoint" return new point instance

	movePointByRef(&p, 1, 1) // change original value by ref
	fmt.Println(p)           // {2 3}

	fmt.Println("---")
	p1 := Point{5, 5}
	p1 = p1.moveOriginalPoint(1, 1)
	fmt.Println(p1) // {6 6}
	fmt.Println("test")
}

type Point struct {
	X, Y int
}

// method for Point structure
func (p Point) moveOriginalPoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

// method by ref
func movePointByRef(p *Point, x, y int) {
	p.X += x
	p.Y += y
}

// method with return new structure instance
func movePoint(p Point, x, y int) Point {
	p.X += x
	p.Y += y
	return p
}

// ------------------------------------------------------------------

// closures (function in function)
func totalPrice(initPrice int) func(int) int {
	sum := initPrice         // run once
	return func(x int) int { // run each time
		sum += x
		return sum
	}
}

// callback
func callbackDoSomething(callback func(int, int) int, s string) int {
	result := callback(5, 2)
	fmt.Println(s)
	return result
}

// ------------------------------------------------------------------

// change method behaviour from outside
func callbackMathExamples() {
	// sum
	sumCallback := func(n1, n2 int) int {
		return n1 + n2
	}
	sumResult := callbackDoSomething(sumCallback, "plus")
	fmt.Println(sumResult) // 7

	// multiple
	multipleCallback := func(n1, n2 int) int {
		return n1 * n2
	}
	multipleResult := callbackDoSomething(multipleCallback, "multiple")
	fmt.Println(multipleResult) // 10
}

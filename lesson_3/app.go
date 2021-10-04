package main

import "fmt"

// Pointers, Structures, Arrays, and Slices
func main() {
	//pointers()
	//structs()
	//arrays()
	//slices()
	//nils()
}

// nil - analog NULL
func nils() {
	var snil []int
	fmt.Println(snil, len(snil), cap(snil)) // [] 0 0
	if snil == nil {
		fmt.Println("slice is nil!")
	}
}

// arrays
func arrays() {
	var a [2]string
	a[0] = "Hello "
	a[1] = "world"
	//a[2] = "!!!" // error - invalid array index
	fmt.Println(a) // [Hello  world]

	numbers := [...]int{1, 2, 3}
	fmt.Println(numbers)

	animals := [4]string{
		"dog",
		"cat",
		"giraffe",
		"elephant",
	}

	// cut slice/array
	animals1 := animals[:2]  // or [0:2]
	animals2 := animals[2:4] // or [2:4]
	animals3 := animals[:]   // or [0:4]
	fmt.Println(animals1)    // [dog cat]
	fmt.Println(animals2)    // [giraffe elephant]
	fmt.Println(animals3)    // [dog cat giraffe elephant]
}

// dynamic length arrays (wrapper over arrays)
func slices() {
	letters := []string{"a", "b", "c"}
	fmt.Println(letters[0])   // a
	fmt.Println(letters)      // [a b c]
	fmt.Println(len(letters)) // 3
	fmt.Println(cap(letters)) // 3

	letters = append(letters, "d", "e")
	fmt.Println(letters)      // [a b c d e]
	fmt.Println(len(letters)) // 5
	fmt.Println(cap(letters)) // 6

	// other way
	createSlice := make([]int, 3)
	fmt.Println(createSlice) // [0 0 0]
}

// link and value by link
func pointers() {
	a := "hello world"
	fmt.Println(a)

	p := &a         // Reference to the memory location of the variable "a"
	fmt.Println(p)  // "0xc000010230" Display the memory location of the variable "a"
	fmt.Println(*p) // Display the contents of the memory location of the variable "a"

	*p = "new value"
	fmt.Println(a)  // new value
	fmt.Println(*p) // new value
}

// ----------------------------------------------------------------------------------------

// Point structure - logical union of a set of fields
type Point struct {
	X int
	Y int
}

// structure method
func (p *Point) someMethod() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
}

// practice with structures
func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
	}

	fmt.Println(p1)
	fmt.Println(p1.X)
	fmt.Println(p1.Y)

	p1.X = 123
	fmt.Println(p1.X)

	// ------------------------

	p2 := Point{X: 123}
	fmt.Println(p2)

	// ------------------------

	g := &p1
	g.Y = 15
	fmt.Println(p1)
	fmt.Println(&g) // "0xc00000e030" Memory address
	fmt.Println(g)  // &{123 15}
	fmt.Println(*g) // {123 15}

	// work with structure methods
	p3 := Point{
		X: 10,
		Y: 20,
	}
	p3.someMethod()
}

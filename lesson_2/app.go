package main

import "fmt"

//const name string = "John"

func main() {
	//surname := "Doe"
	//fmt.Println(fmt.Sprintf("%s %s", name, surname))

	// Will be executed at the end of the script LIFO (Last In, First Out)
	defer fmt.Println("Delayed operation 1")
	defer fmt.Println("Delayed operation 2")
	defer fmt.Println("Delayed operation 3")

	//fmt.Println(forLoop(10))
	conditionsTest()
}

// struct instead class
type testStruct struct {
	a string
	b string
	c string
}

// func with return type "struct"
func test3WithStruct() testStruct {
	stringStruct := testStruct{
		a: "Hello",
		b: " all ",
		c: "world",
	}

	return stringStruct
}

// basic func with return multi param types
func test() (string, string) {
	a := "hello"
	b := "world"

	return a, b
}

// basic func with return one param types
func test1() string {
	a := "hello"
	return a
}

// Out parameters have already been initialized
func namedReturnParams() (a, b string) {
	a = "Hello "
	b = "world"
	return a, b // can be "return" but not recommend.
}

// simple "for" loop
func forLoop(maxVal int) int {
	sum := 0
	for i := 0; i < maxVal; i++ {
		sum += i
	}

	return sum
}

// analog "while" loop
func whileLoop(maxVal int) {
	sum := 0
	for sum < maxVal {
		sum += 1
	}

	fmt.Println(sum)
}

// endless loop
func endlessLoop() {
	sum := 0
	for {
		sum += 1
		fmt.Println(sum)
	}
}

// simple if and if with initialize, switch
func conditionsTest() {
	desiredResult := 45
	for i := 0; i <= 10; i++ {
		if j := forLoop(i); j == desiredResult {
			fmt.Println(desiredResult)
		}

		if i == 3 {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 10; i++ {
		switch j := forLoop(i); j {
		case desiredResult:
			fmt.Println(desiredResult)
		default:
			fmt.Println(i)
		}
	}
}

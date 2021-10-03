package main

import "fmt"

const name string = "John"

func main() {
	surname := "Doe"
	fmt.Println(fmt.Sprintf("%s %s", name, surname))
}

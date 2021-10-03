package main

//const name string = "John"

func main() {
	//surname := "Doe"
	//fmt.Println(fmt.Sprintf("%s %s", name, surname))
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

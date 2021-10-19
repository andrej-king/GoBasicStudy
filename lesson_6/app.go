package main

import "fmt"

// interfaces
func main() {
	//var a someInterface
	//ss := someStructure{1, 2}
	//os := otherStructure{3, 4}
	//
	//a = &ss
	//fmt.Println(a.sum()) // 3
	//a = os
	//fmt.Println(a.sum()) // 7
	//
	//var i someInterface = otherStructure{5, 10}
	//fmt.Println(i.sum()) // 15

	//fmt.Printf("%v, %T\n", i, i) // {5 10}, main.otherStructure

	//var os1 *someStructure
	//var emptyInterface someInterface
	//fmt.Printf("%v, %T\n", emptyInterface, emptyInterface) // <nil>, <nil>
	//emptyInterface = os1
	//fmt.Printf("%v, %T\n", emptyInterface, emptyInterface) // <nil>, *main.someStructure

	// Casting to type
	//var a interface{} = "hello"
	//s := a.(string) // Casting to string type
	//fmt.Println(s)

	// Check if it is possible to cast to this type
	//s, ok := a.(string)
	//fmt.Println(s, ok) // hello true

	//s, ok := a.(float32)
	//fmt.Println(s, ok) // 0 false
	//if ok {
	//	//write logic for use var "s"
	//}

	//f := a.(float32)
	//fmt.Println(f) // panic: interface conversion: interface {} is string, not float32

	// check var type
	var b interface{} = true
	switch b.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Printf("unknown type %T\n", b) // unknown type bool
	}

	// structure for JSON
	//var json map[string]interface{} // string key, value any type

}

type someInterface interface {
	sum() int
}

type someStructure struct {
	N1, N2 int
}

// someStructure realise for someInterface
func (s *someStructure) sum() int {
	return s.N1 + s.N2
}

type otherStructure struct {
	A, B int
}

// otherStructure realise for someInterface
func (o otherStructure) sum() int {
	return o.A + o.B
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

//env for linux: GOOS=linux;GOARCH=amd64
//env for windows: GOOS=windows;GOARCH=amd64
//env for macos: GOOS=darwin;GOARCH=amd64

// build
//extract GOROOT=D:\Go\go1.17.1 #gosetup
//extract GOPATH=D:\Go\go1.17.1\bin #gosetup
//D:\Go\go1.17.1\bin\go.exe build -o D:\GoProjects\syntax\build\lesson_1\app_windows.exe D:\GoProjects\syntax\lesson_1\app.go #gosetup

// var, types, concat, print, run simple HTTP server
func main() {
	//var a int8 = -1
	//var b int16 = -1
	//var c int32 = -1
	//var d int64 = -1
	//var e uint8 = 1
	//var f uint16 = 1
	//var g uint32 = 1
	//var h uint64 = 1
	//var i byte = 1 // uint8
	//var j rune = -1 // int32
	//var k int = -1 // int32 or int 64
	//var m uint = 1 // uint32 or uint64

	//var a1 float32 = 31.11
	//var a2 float64 = 64.11

	//var b1 bool = true
	//var b2 bool = false

	var name string = "Hello world"
	fmt.Println(name)

	var name3 string = "John" // equals name4 and name5
	var name4 = "John"        // equals name3 and name5
	name5 := "John"           // equals name3 and name4
	fmt.Println(name3)
	fmt.Println(name4)
	fmt.Println(name5)

	var (
		name6 = "John"
		age   = 20
	)
	c := fmt.Sprintf("Hello %s. You are %d years old.", name6, age) // fmt.Sprintf concat
	fmt.Println(c)
	fmt.Println(name5)

	// run simple web server
	runServer()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!\n"))
}

func runServer() {
	http.HandleFunc("/", Handler)

	log.Println("Start HTTP server on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

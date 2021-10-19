package main

import (
	"fmt"
	"io"
	"strings"
)

// Reader interface (Reading the stream in bytes.)
func main() {
	//io.Reader()
	//io.EOF // ond of file
	r := strings.NewReader("hello world")

	arr := make([]byte, 8) // read by 8 bytes

	// endless loop. Reading the stream in bytes.
	for {
		n, err := r.Read(arr) // Read write data in "arr" slice
		fmt.Printf("n = %d err = %v arr = %v\n", n, err, arr)
		fmt.Printf("arr n bytes %q", arr[:n])
		if err == io.EOF {
			break
		}
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// work with files and run shell commands
func main() {
	//readFile()
	//writeToFile()
	//appendToFile()

	shellCommand()
}

// shellCommand run shell commands
func shellCommand() {
	//c := exec.Command("ssh", "-t", "localhost", "top") // command name and param counter by delimiter ","
	c := exec.Command("cat", "./lesson_13/test.txt") // command name and param counter by delimiter ","
	//c := exec.Command("top") // command name
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}

// appendToFile append data in file (or create new if note exists).
func appendToFile() {
	file, err := os.OpenFile("./lesson_13/test.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	if _, err = file.WriteString(" append some data"); err != nil {
		panic(err)
	}
}

// writeToFile override data in file.
func writeToFile() {
	data := []byte("Override data")
	err := ioutil.WriteFile("./lesson_13/test.txt", data, 0600)
	if err != nil {
		panic(err)
	}
}

// readFile print data from file.
func readFile() {
	data, err := ioutil.ReadFile("./lesson_13/test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

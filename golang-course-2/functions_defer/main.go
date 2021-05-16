package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo")
}

func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is foobar")
}

func main() {
	defer foo()
	bar()

	defer foobar()
	bar()

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// code that works with file
}

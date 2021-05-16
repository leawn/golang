package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("Anonymous function") // call it immediately cause it has no name to call later

	a := increment(10)
	a()
	fmt.Println(a())
}

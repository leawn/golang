package main

import "fmt"

func main() {
	var i int

loop:
	if i < 3 {
		i++
		goto loop
	}
	fmt.Println("done")
}

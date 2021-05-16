package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v\n", p1, &p1)
	fmt.Printf("Value of pp1: %v, address of pp1: %v\n", pp1, &pp1)

	fmt.Println(*p1)
	fmt.Println(**pp1) // actually using two stars means pointer from pointer

	// comparing pointers
	var p2 *int
	fmt.Println(p2)

	y := 5
	p2 = &y
	z := 5
	p3 := &z

	fmt.Println(p2 == p3)

	// pointers are equal when they point to the same value (cause you compare the address of where the variable is actually stored, so ye)
} // but if you compare the values stores using *, then it's alright, it just compares the values

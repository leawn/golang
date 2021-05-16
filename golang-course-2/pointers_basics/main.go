package main

import "fmt"

func main() {
	name := "Andrei"
	fmt.Println(&name)

	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type of %T with a value of %v and address %p\n", ptr, ptr, &ptr)
	fmt.Println(ptr)

	var ptr1 *float64 // zero value is nil
	_ = ptr1

	p := new(int)

	x = 100
	p = &x

	fmt.Printf("p is of type %T with a value of %v\n", p, p)

	*p = 90 // equivalent to x = 90 or p = &x | *p = x
	fmt.Println(x, *p)
}

package main

import "fmt"

func main() {
	var x = 3   // int type
	var y = 3.1 // float64 type
	//x = x * y
	x = x * int(y)
	fmt.Println(x)
	fmt.Println(y)

	var a = 5 // int type
	fmt.Printf("%T\n", a)

	var b int64 = 2
	_ = b
}

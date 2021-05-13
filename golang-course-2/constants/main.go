package main

import "fmt"

func main() {
	const days int = 7

	var i int
	_ = i

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234 //in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	x, y := 5, 1
	fmt.Println(x / y) // run-time error

	const a = 5
	const b = 0
	//fmt.Println(a / b) // compile error

	const n, m int = 4, 5
	const n1, m1 = 3, 6

	const (
		min1 = -500
		min2 // equals -500 too, cause it depicts the previous constant
		min3 = 100
	)

	fmt.Println(min1, min2, min3)
}

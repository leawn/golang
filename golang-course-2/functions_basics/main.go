package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1")
}

func f2(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	s = a + b
	return
}

func main() { // golang automatically executes only "main" and "init" functions
	f1()
	f2(5, 7)
	f3(1, 2, 3, 4.3, 5.3, "nice")
	p := f4(3.4)
	fmt.Println(p)
	x, y := f5(1, 2)
	fmt.Println(x, y)
	l := sum(1, 2)
	fmt.Println(l)
}

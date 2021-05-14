package main

import "fmt"

func main() {
	numbers := []int{2, 3}

	numbers = append(numbers, 10)
	fmt.Println(numbers)

	numbers = append(numbers, 1, 2, 3, 4, 5)
	fmt.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	src := []int{10, 20, 30}
	dst := make([]int, len(src)) // making a new slice with required length
	fmt.Println(dst)
	nn := copy(dst, src) // copying the contents
	fmt.Println(src, dst, nn)
}

package main

import "fmt"

type emptyInterface interface{}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "golang"
	fmt.Println(empty)

	empty = []int{3, 4, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int))) // type assertion
	empty = "gl"
	fmt.Println(len(empty.(string))) // alright?

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []int{1, 2, 3, 7}
	fmt.Println(you.info)
}

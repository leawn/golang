package main

import "fmt"

func main() {
	s1 := "I love golang"
	fmt.Println(s1[2:5])

	rs := []rune(s1) // better to convert all the strings to slices of runes and the slice them again, that way we'll show all the characters in the right way
	fmt.Printf("%T\n", rs)

	fmt.Println(string(rs[0:3]))
}

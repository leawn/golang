package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	//a[start:stop] the scheme of slice
	// start is included, stop is excluded
	b := a[1:3]
	fmt.Println(b)
	fmt.Printf("%T\n", b) // b is of type slice

	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:4]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // when omitted uses the default value (last element)
	fmt.Println(s1[:3]) // the same default value (first element)
	fmt.Println(s1[:])  // first and last element

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)
}

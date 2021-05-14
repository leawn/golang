package main

import "fmt"

func main() {
	names := [5]string{"Helen", "Mark", "Antonio", "Michael", "Marta"}
	friends := [2]string{"Mark", "Marry"}

outer:
	for index, name := range names {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")
}

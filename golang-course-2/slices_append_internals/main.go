package main

import "fmt"

func main() {
	var nums []int
	fmt.Printf("%#v\n", nums)

	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 1, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	nums = append(nums, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums), cap(nums))

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[:1], "X", "Y")
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters))

	//fmt.Println(letters[4]) // gives an error
	fmt.Println(letters[3:6])
}

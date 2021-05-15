package main

import "fmt"

func main() {
	friends := map[string]int{"Dan": 40, "Maria": 25}

	neighbors := friends

	friends["Dan"] = 50
	fmt.Println(neighbors)

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}

	fmt.Println(people)
	people["Anne"] = 80
	fmt.Println(people)
	fmt.Println(friends)
}

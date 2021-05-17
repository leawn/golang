package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	c := make(chan int)

	// only for receiving || unidirectional channel
	c1 := make(<-chan string)

	// only for sending
	c2 := make(chan<- chan string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	n := <-c
	fmt.Println(n)
}

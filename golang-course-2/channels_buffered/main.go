package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // buffered channel

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d ends sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c { // v := <- c
		fmt.Println("main goroutine received value from channel:", v)
	}

	fmt.Println(<-c)

	// c <- 10 // panic: send on a closed channel

}

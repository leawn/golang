package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down!", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		c <- url
	} else {
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is up!\n", url)
		fmt.Println(s)
		c <- url

	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	// 1.
	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkUrl(<-c, c)
	// 	fmt.Println(strings.Repeat("&", 30))
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range c {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, c)
	// }

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url)
	}
}

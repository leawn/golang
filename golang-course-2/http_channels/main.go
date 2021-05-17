package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is down!", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s += fmt.Sprintf("Error: %v\n", err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"

			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				s += fmt.Sprintf("Error: %v\n", err)
				c <- s
			}
			s += fmt.Sprintf("%s is up!\n", url)
			c <- s
		}
	}
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	// 1.
	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

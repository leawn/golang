package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args", os.Args)
	fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])
	fmt.Println("Number of arguments:", len(os.Args)-1)

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(result)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err
}

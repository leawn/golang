package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b'

	fmt.Printf("Type %T, Value: %d\n", var1, var1)
	fmt.Printf("Type %T, Value: %d\n", var2, var2)

	str := "tar"
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c\n", r)
		i += size
	}
	fmt.Printf("\n")

	for _, r := range str {
		fmt.Printf("%c\n", r)
	}
}

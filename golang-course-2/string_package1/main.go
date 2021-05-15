package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	result := strings.Contains("I love golang", "love")
	p(result)

	result = strings.ContainsAny("success", "sc")
	p(result)

	result = strings.ContainsRune("golang", 'm')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "")
	p(n)

	p(strings.ToLower("GO JAVA SCALA"))
	p(strings.ToUpper("kafka flint spark"))
	p("go" == "gO")

	p(strings.EqualFold("GO", "go"))

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr)

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr)

	s := strings.Split("a,b,c", ",")
	p(s)

	s = []string{"I", "learn", "golang"}
	myStr = strings.Join(s, "")
	p(myStr)

	myStr = "Orange Green \n Blue Yellow"
	fields := strings.Fields(myStr)
	fmt.Println(fields)
	fmt.Printf("%T\n", fields)
	fmt.Printf("%#v\n", fields)

	s1 := strings.TrimSpace("\t hey  there\n")
	p(s1)

	s2 := strings.Trim("...Hello gophers!!!!?", ".!?")
	p(s2)
}

package main

import f "fmt" // we can rename imports

const done = false // package scoped

var b int = 10 // if variable is package scoped it must not be necessarily used (cause it be used in other packages that import this one, for example)

func main() {
	var task = "Running" // local(block) scoped
	_ = task

	f.Println("ez")

	f.Println("done equals:", done)
	const done = true // local(block) scoped
	f.Println("done equals:", done)

	f1() // anyway package scope is used
}

func f1() {
	f.Println("done in f1():", done)                  // package scope again
	const done = true                                 // but if we redeclare it
	f.Println("done in f1() after redeclaring", done) // then it equals that redeclaration
}

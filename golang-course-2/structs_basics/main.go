package main

import "fmt"

func main() {
	title1, author1, year1 := "The Divine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Effective TypeScript", "Dan Vaderkam", 2019

	fmt.Println("Book1:", title1, author1, year1)
	fmt.Println("Book2:", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"Effective TypeScript", "Dan Vaderkam", 2019}
	fmt.Println(myBook)

	bestBook := book{title: "SomeTitle", author: "SomeAuthor", year: 1000}
	_ = bestBook

	aBook := book{title: "A random book"}
	fmt.Println(aBook)
}

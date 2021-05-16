package main

import "fmt"

func main() {
	type book struct {
		title  string
		author string
		year   int
	}

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	lastBook.author = "Lev Tolstoy"
	lastBook.year = 1878
	fmt.Println(lastBook)

	aBook := book{"Anna Karenina", "Lev Tolstoy", 1878}

	fmt.Println(aBook == lastBook)

	myBook := aBook
	myBook.year = 2021
	fmt.Println(myBook, "\n", lastBook)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       20,
	}

	fmt.Printf("%#v\n", diana)
	fmt.Println("Diana's age:", diana.age)

	type Book struct { // anonymous struct
		string
		float64
		bool
	}

	b1 := Book{"1984 by George Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 40000, false}
	fmt.Printf("%#v\n", e)
}

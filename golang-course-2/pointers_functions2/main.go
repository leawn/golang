package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.3
	name = "New"
	sold = false
}

func changeValuesPointers(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.3
	*name = "New"
	*sold = false
}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bike"
}

func changeProductPointers(p *Product) {
	(*p).price = 300
	(*p).productName = "Bike"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 1, 200.2, "Old", true
	fmt.Println("Before calling:", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After calling:", quantity, price, name, sold) // they stay the same, cause golang passes by copies of values, not by the references
	changeValuesPointers(&quantity, &price, &name, &sold)
	fmt.Println("After calling with pointers", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Smth",
	}
	changeProduct(gift)
	fmt.Println(gift)
	changeProductPointers(&gift)
	fmt.Println(gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	alphabet := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	changeMap(alphabet)
	fmt.Println(alphabet)

	// so for working and changing bools, ints, strings, structs we should pass a pointer to a value
	// however for working with slices or maps you shouldn't pass any pointers cause maps and slices save the reference, not the value itself
}

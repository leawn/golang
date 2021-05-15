package main

import "fmt"

func main() {
	var employees map[string]string // uninitialized
	fmt.Printf("%#v\n", employees)
	fmt.Println(employees)

	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1

	//employees["Dan"] = "Programmer"

	people := map[string]float64{} // initialized
	fmt.Println("########")
	fmt.Println(people)
	people["John"] = 21.4
	people["Marry"] = 10
	fmt.Println(people)

	map1 := make(map[int]int)
	fmt.Println(map1)
	map1[4] = 8
	fmt.Println(map1)

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 121.12,
		"RON": 213.2,
	}

	balances["USD"] = 500.2
	balances["GBP"] = 100
	fmt.Println(balances)

	v, ok := balances["RON"] // in "ok" we get true/false value
	fmt.Println(v, ok)

	if ok {
		fmt.Println("The RON balance is:", v)
	} else {
		fmt.Println("The RON key does not exist in the map")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	delete(balances, "USD")
	fmt.Println(balances)
}

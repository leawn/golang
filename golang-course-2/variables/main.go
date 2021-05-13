package main

import "fmt"

func main() {
	var age int = 20
	fmt.Println("Age:", age)

	var name = "Dan"
	//fmt.Println("Your name is:", name)
	_ = name

	s := "Learning golang!"
	_ = s

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var i, j int
	i, j = 5, 8

	j, i = i, j

	fmt.Println(i, j)

	var a = 4
	var b = 4.5

	a = int(b)
	fmt.Println(a, b)

	figure := "Circle"
	radius := 5
	pi := 3.141590

	fmt.Printf("Radius is %d\n", radius) // %d for a decimal (base10)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi) // %f for a decimal point but no exponent

	// %s is for a string
	fmt.Printf("The diameter of a %s with a radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q\n", figure) // %q for a quoted string

	// %v for any value

	// %T for dispaying a type of a variable
	fmt.Printf("The type of %v is %T\n", radius, radius)

	// %t for boolean
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b for base2
	fmt.Printf("%08b \n", 55)

	// we can say how many of numbers we wanna see after the floating point
	x := 3.2
	y := 5.7
	fmt.Printf("x * y = %.3f\n", x*y)
}

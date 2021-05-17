package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// func printCircle(c circle) {
// 	fmt.Println("Shape:", c)
// 	fmt.Println("Area:", c.area())
// 	fmt.Println("Perimeter:", c.perimeter())
// }

// func printRectangle(r rectangle) {
// 	fmt.Println("Shape:", r)
// 	fmt.Println("Area:", r.area())
// 	fmt.Println("Perimeter:", r.perimeter())
// }

func print(s shape) {
	fmt.Println("Shape:", s)
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	// var s shape
	// fmt.Printf("%T\n", s)

	// ball := circle{radius: 2.5}
	// s = ball

	// print(s)
	// fmt.Printf("%T\n", s)

	// room := rectangle{width: 2, height: 3}
	// s = room
	// fmt.Printf("%T\n", s)

	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	//s.volume() shape interface has no volume() method implemented
	ball, ok := s.(circle) // assertion can fail

	if ok {
		fmt.Println("Ball volume:", ball.volume())
	}

	s = rectangle{width: 3.4, height: 2.}
	switch value := s.(type) {
	case circle:
		fmt.Println(value, "has circle type")
	case rectangle:
		fmt.Println(value, "has rectangle type")
	}
}

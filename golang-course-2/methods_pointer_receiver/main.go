package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

type distance *int

//func (d *distance) m1() { you can not implement methods to pointers
//	fmt.Println()
//}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "bmw", 35000)
	fmt.Println(myCar)

	myCar.changeCar1("porsche", 80000)
	fmt.Println(myCar)

	myCar.changeCar2("MB", 42000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	yourCar.changeCar2("nissan", 20000) // you can do this way
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("vw", 23000) // or this way (golang interprets if a receiver function needs a pointer)
	fmt.Println(*yourCar)
}

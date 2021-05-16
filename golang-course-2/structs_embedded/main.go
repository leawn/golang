package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Keller",
		salary: 3000,
		contactInfo: Contact{
			email:   "jkeller@company.com",
			address: "Street 20, London",
			phone:   001222234,
		},
	}

	fmt.Printf("%+v\n", john)

	fmt.Println("Employee's email:", john.contactInfo.email)

	john.contactInfo.email = "new_email@company.com"

	fmt.Println("Employee's email:", john.contactInfo.email)
}

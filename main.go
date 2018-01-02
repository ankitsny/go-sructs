package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pPtr *person) updateName(uName string) {
	(*pPtr).firstName = uName
}

func main() {
	// Method #1
	user1 := person{"Virat", "Kohli", contactInfo{"ankit@email.com", 123}}

	// Method #2
	user2 := person{
		firstName: "John",
		lastName:  "Doe",
		contact: contactInfo{
			email:   "ankso@ankso.com",
			zipCode: 12345,
		},
	}

	// Method #3
	var user3 person
	// fmt.Printf("%+v", user3) // display props name
	user3.firstName = "FirstNAME"
	user3.lastName = "LastNAME"
	user3.contact.email = "ankso@ankso.com"
	user3.contact.zipCode = 12322

	// Method #1
	(&user1).updateName("ANKSO")
	user1.print()

	// Method #2 - Shortcut
	// Automatic type conversion
	user2.updateName("HelloName")

	user2.print()
	user3.print()
}

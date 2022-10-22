package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := Person{"Alex", "Anderson"}
	// fmt.Println(alex)

	// var alex Person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	var jim Person
	jim.firstName = "Jim"
	jim.lastName = "Party"
	jim.contact = ContactInfo{
		email:   "jim@gmail.com",
		zipCode: 9600,
	}
	// pointer
	// jimPointer := &jim

	jim.updateName("Huy")
	jim.print()

}

func (pointerToPerson Person) print() {
	fmt.Printf("%+v", pointerToPerson)
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

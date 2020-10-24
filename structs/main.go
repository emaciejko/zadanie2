package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Eryk")
	jim.print()
}

func (pointerToPerson *person) updateName(n string) {
	(*pointerToPerson).firstName = n
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

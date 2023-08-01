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
		lastName:  "Doe",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 95110,
		},
	}

	jim.updateName("Rahul")
	jim.print()
}

func (pTp *person) updateName(newFName string) {
	(*pTp).firstName = newFName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

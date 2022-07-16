package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

func main() {
	// alex := person{"Alex", "anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}  // same as the above line but more dynamic
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anders"
	// fmt.Println(alex)
	// fmt.Printf(" %+ v", alex)
	jim := person{
		firstName: "Jim",
		lastName: "Gyms",
		contactInfo: contactInfo{
			email: "me@example.com",
			zipCode: 87777,},
	}
		jimPointer := &jim
		jimPointer.updateName("James")
		jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

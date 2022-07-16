package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName string
	contact contactInfo
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
		contact: contactInfo{
			email: "me@example.com",
			zipCode: 87777,},
	}
		fmt.Printf("%+v", jim)
}

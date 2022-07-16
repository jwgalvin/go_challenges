package main

import "fmt"
type person struct {
	firstName string
	lastName string
}

func main() {
	// alex := person{"Alex", "anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}  // same as the above line but more dynamic
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anders"
	fmt.Println(alex)
	fmt.Printf(" %+ v", alex)
}
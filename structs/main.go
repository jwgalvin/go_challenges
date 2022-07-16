package main

import "fmt"
type person struct {
	firstName string
	lastName string
}

func main() {
	// alex := person{"Alex", "anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}  // same as the above line but more dynamic
	
	fmt.Println(alex)
}
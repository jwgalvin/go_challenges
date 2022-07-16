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
	jim := person{
		firstName: "Jim",
		lastName: "Gyms",
		contact: contactInfo{
			email: "me@example.com",
			zipCode: 87777,},
	}
		fmt.Printf("%+v", jim)
}

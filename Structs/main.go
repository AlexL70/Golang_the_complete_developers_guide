package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email: "jim@somemail.com",
			zip:   9400,
		},
	}

	fmt.Printf("%+v", alex)
}

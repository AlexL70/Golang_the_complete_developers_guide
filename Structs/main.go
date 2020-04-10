package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{
	// 	firstName: "Alex",
	// 	lastName: "Levinson"}
	// fmt.Println(alex)

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Levinson"
	fmt.Printf("%+v", alex)
}

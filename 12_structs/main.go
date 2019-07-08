package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "f" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Ferdous",
		lastName: "Haque",
		city:     "Dhaka",
		gender:   "M",
		age:      25}

	// Alternative
	person2 := Person{"Tonni",
		"Haque",
		"Dhaka",
		"M",
		25}

	/* fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1) */

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Test")
	person2.getMarried("Test22")
	fmt.Println(person2.greet())

}

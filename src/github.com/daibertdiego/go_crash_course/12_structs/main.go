package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Method value receiver (doesn't change anything)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Method pointer receiver (changes something)
func (p *Person) hasBirthday() {
	p.age++
}

// Method pointer receiver (changes something)
func (p *Person) getMarried(spouseLastName string) {
	if strings.EqualFold(p.gender, "female") {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person
	person := Person{
		firstName: "Diego",
		lastName:  "Almeida",
		city:      "Rochester",
		gender:    "Male",
		age:       35,

		// firstName, lastName, city, gender string
		// age int
	}

	// Alternative way
	person2 := Person{
		"Raquel",
		"Neves",
		"Rochester",
		"female",
		35,
	}

	fmt.Println(person, person2)
	fmt.Println(person.firstName, person2.firstName)
	person.age++
	fmt.Println(person.age)

	fmt.Println(person.greet())
	person.hasBirthday()
	person.hasBirthday()
	person.hasBirthday()
	person.hasBirthday()
	fmt.Println(person.greet())
	person.getMarried("Daibert")
	person2.getMarried("Daibert")
	fmt.Println(person.greet())
	fmt.Println(person2.greet())
}

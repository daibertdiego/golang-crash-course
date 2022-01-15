package main

import "fmt"

var globalVariable = true

// invalidDeclaration := true syntax error: non-declaration statement outside function body

func main() {
	// Main Types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alias for uint8
	// rune - alias for uint32
	// float32 float64
	// complex64 complex128

	var name string = "Galão da Massa"
	var name2 = "Galão da Massa"
	// Shorthand
	name3 := "Galão da Massa"

	fmt.Println(name, name2, name3)

	var age = 35
	fmt.Printf("%v %T \n", age, age)

	const ageConst = 13
	// ageConst = 15 Forbiden
	fmt.Printf("%v %T \n", ageConst, ageConst)

	fmt.Printf("%v %T \n", globalVariable, globalVariable)

	size := 1.3
	var size2 float32 = 1.3
	fmt.Printf("%v %T \n", size, size)
	fmt.Printf("%v %T \n", size2, size2)

	name, email := "Diego", "daibertdiego@gmail.com"
	fmt.Printf("%v %v %T \n", name, email, name)
}

package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assing values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr[0], fruitArr[1])

	// Declare and assign
	fruitArrAssigned := [2]string{"Grape", "Coconut"}
	fmt.Println(fruitArrAssigned)

	// Slice
	fruitSlice := []string{"Apple", "Orange", "Grape", "Coconut", "Watermelon"}
	fmt.Println("List size", len(fruitArr), fruitSlice)
	fmt.Println("Only index 2 and 3", fruitSlice[1:3])
}

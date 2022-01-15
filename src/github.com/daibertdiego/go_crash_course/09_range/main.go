package main

import "fmt"

func main() {
	ids := []int{1, 2, 345, 65, 44, 35, 67, 13}

	for i, id := range ids {
		fmt.Printf("Index: %d - Id: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("Id: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	soccer_roster := map[string]int{"Diego": 9, "Alan": 11, "Ramon": 2}
	for k, v := range soccer_roster {
		fmt.Printf("%s: %d \n", k, v)
	}

	for k := range soccer_roster {
		fmt.Printf("%s \n", k)
	}
}

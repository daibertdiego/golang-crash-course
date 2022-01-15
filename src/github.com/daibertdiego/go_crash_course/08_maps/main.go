package main

import "fmt"

func main() {
	// Define maps
	emails := make(map[string]string)

	// Assign key-value
	emails["Diego"] = "daibertdiego@gmail.com"
	emails["Raquel"] = "raquelalneves@gmail.com"

	fmt.Println(emails, "Size", len(emails))

	for k, v := range emails {
		fmt.Println(k, v)
	}

	delete(emails, "Diego")
	fmt.Println(emails)

	// Define and Assign maps
	emailsMap := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(emailsMap)
}

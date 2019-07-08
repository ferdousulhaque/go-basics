package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// emails["Ferdous"] = "ferdousul.haque@gmail.com"
	// emails["Tonni"] = "toimoi@gmail.com"
	// emails["Rakib"] = "rakib@vromordigital.com"
	emails := map[string]string{"Bob": "bob@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Ferdous"])

	delete(emails, "Rakib")
	fmt.Println(emails)
}

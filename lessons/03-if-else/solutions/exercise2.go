package main

import "fmt"

func main() {

	var role string

	fmt.Print("Role: ")
	fmt.Scanln(&role)

	if role == "admin" {
		fmt.Println("Full access granted.")
	} else {
		fmt.Println("Limited access.")
	}
}
package main

import "fmt"

func main() {

	var env string

	fmt.Print("Enter environment: ")
	fmt.Scanln(&env)

	switch env {
	case "development":
		fmt.Println("Development environment")

	case "staging":
		fmt.Println("Staging environment")

	case "production":
		fmt.Println("Production environment")
	default:
		fmt.Println("Unknown environment.")
	}
}

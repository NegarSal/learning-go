package main

import "fmt"

func main() {

	var environment string
	var portNumber int

	fmt.Print("Enter environment: ")
    fmt.Scanln(&environment)

	fmt.Print("Enter port number: ")
    fmt.Scanln(&portNumber)

	fmt.Println()
	fmt.Println("Environment: ", environment)
	fmt.Println("Port Number: ", portNumber)

}

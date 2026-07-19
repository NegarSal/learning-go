package main

import "fmt"

func main() {

	var serverName string
	var environment string

	fmt.Print("Enter server name: ")
	fmt.Scanln(&serverName)

	fmt.Print("Enter environment: ")
	fmt.Scanln(&environment)

	fmt.Println()
	fmt.Println("Server Name:", serverName)
	fmt.Println("Environment:", environment)
}
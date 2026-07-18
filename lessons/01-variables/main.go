package main

import "fmt"

func main() {
	serverName := "web-01"
	environment := "production"

	fmt.Println("Server:", serverName)
	fmt.Println("Environment:", environment)

	serverName = "web-02"

	fmt.Println("Updated Server:", serverName)
}
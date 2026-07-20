package main

import "fmt"

func main() {

	var port int

	fmt.Print("Enter the port number: ")
	fmt.Scanln(&port)

	if port < 1 {
		fmt.Println("Port number cannot be less than 1.")
	} else if port <= 1023 {
		fmt.Println("Well-known Port")
	} else if port <= 65535 {
		fmt.Println("Valid User Port")
	} else {
		fmt.Println("Port number cannot be greater than 65535.")
	}
}

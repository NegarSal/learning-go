package main

import "fmt"

func main() {

	var status string
    
	fmt.Print("Enter Server Status: ")
	fmt.Scanln(&status)

	if status == "online" {
		fmt.Println("Server is online.")
	} else {
		fmt.Println("Server is offline.")
	}
}
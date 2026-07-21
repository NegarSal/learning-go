package main

import "fmt"

func main() {

	var status string

	fmt.Print("Enter status: ")
	fmt.Scanln(&status)

	switch status {
	case "running":
		fmt.Println("Container is running.")
	case "stopped":
		fmt.Println("Container is stopped.")
	case "paused":
		fmt.Println("Container is paused.")
	default:
		fmt.Println("Unknown status.")
	}
}

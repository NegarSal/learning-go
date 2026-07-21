package main

import "fmt"

func main() {

	var service string

	fmt.Print("Enter service name: ")
	fmt.Scanln(&service)

	switch service {
	case "nginx":
		fmt.Println("Web Server")
	case "mysql":
		fmt.Println("Database Server")
	case "redis":
		fmt.Println("Cache Server")
	default:
		fmt.Println("Unknown service.")
	}
}

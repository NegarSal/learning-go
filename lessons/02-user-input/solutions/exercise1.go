package main

import "fmt"

func main() {

	var containerName string
	var imageName string

	fmt.Print("Enter Container Name: ")
    fmt.Scanln(&containerName)

	fmt.Print("Enter Image Name: ")
    fmt.Scanln(&imageName)

	fmt.Println()
	fmt.Println("Container Name: ", containerName)
	fmt.Println("Image Name: ", imageName)

}

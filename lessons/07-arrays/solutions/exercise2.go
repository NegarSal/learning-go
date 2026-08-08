package main

import "fmt"

func main() {
	colors := [...]string{"Red", "Green", "Blue","Yellow"}

	fmt.Println("Colors:", colors)

	fmt.Println("Iterating over colors:")
	for index, value:= range colors{
		fmt.Println("index:", index ," value:", value)
	}
}
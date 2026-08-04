package main

import "fmt"

func isEven(a int) bool {
	return a%2 == 0
}

func main() {
	fmt.Println(isEven(6))
	fmt.Println(isEven(5))
}
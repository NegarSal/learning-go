package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "grape"
	fruits[3] = "melon"

	fmt.Println(fruits)
	fmt.Println("Length:", len(fruits))

	fruits[1] = "kiwi"
	fmt.Println(fruits)
}

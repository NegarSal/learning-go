package main

import "fmt"

func main() {
	var scores [3]int

	scores[0] = 90
	scores[1] = 85
	scores[2] = 100

	fmt.Println(scores)
	fmt.Println("Length:", len(scores))
}

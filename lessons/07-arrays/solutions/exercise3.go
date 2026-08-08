package main

import "fmt"

func main() {

	// Define a two-dimensional array (2 rows and 3 columns)
	var matrix [2][3]int

	// Assign values to the first row
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	// Assign values to the second row
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6

	// Print the two-dimensional array using nested loops
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println() // Move to the next line after printing each row
	}
}

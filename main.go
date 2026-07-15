package main

import "fmt"

func main() {

  var choice int

  fmt.Print("Enter the choice:")
  fmt.Scanln(&choice)

  switch (choice){
  case 1:
	     fmt.Println("Linux")
  case 2:
	     fmt.Println("Docker")
  case 3:
	     fmt.Println("Kubernetes")
  default:
	     fmt.Println("Unknown")	

  }
}

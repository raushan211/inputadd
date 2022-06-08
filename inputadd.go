package main

import "fmt"

// main function
func main() {

	// Println function is used to
	// display output in the next line
	fmt.Println("Enter Your First Number: ")

	// var then variable name then variable type
	var first int

	// Taking input from user
	fmt.Scanln(&first)
	fmt.Println("Enter Second Number: ")
	var second int
	fmt.Scanln(&second)

	// Print function is used to
	// display output in the same line
	fmt.Print("Addition: ")

	// Addition of two int
	fmt.Print(first + " " + second)
}

package main

import "fmt"

func main() {
	// --------------------------------------------
	// Declare variable
	var name string = "John Doe"
	var age int = 10
	var isMarried bool = true

	// --------------------------------------------
	// Print variable
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	// --------------------------------------------
	// Print type of variable
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)
}
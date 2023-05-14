package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)  // <nil>

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer", ptr)
	fmt.Println("Value of actual pointer", *ptr)

	*ptr = *ptr + 1
	fmt.Println("new value is", myNumber)
}

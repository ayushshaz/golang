package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	// No inheritance, No super or parent
	ayush := User{"Ayush","email@gmail.com"}
	fmt.Println(ayush)

	// if else

	if num:= 3; num<10 {
		fmt.Println("nums is less than 10")
	}
}

type User struct{
	Name string
	Email string
}

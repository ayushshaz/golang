package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions and defers")

	proAns, message := add(1,2,3);
	fmt.Println("Addition is", proAns,"and message is", message)

// defer 
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("three")
	fmt.Println("Four")

	// All defer part comes to like a stack and from top of stack each one gets executed
}

func add(values ...int) (int,string){
	ans := 0
	for _,value := range values{
		ans+=value
	}
	return ans,"Error"
}

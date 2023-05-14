package main

import "fmt"

// Capital L is for public
const Logintoken string = "token" 

func main() {
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variables is of type : %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
 	fmt.Printf("Variables is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variables is of type : %T \n", smallVal)

	var smallfloat float32 = 234.23145134534154351 
	fmt.Println(smallfloat)
	fmt.Printf("Variables is of type : %T \n", smallfloat)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is of type : %T \n", anotherVariable)

	// implicit type
	// lexer by default adds the data type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style || walrus operator
	// this scope is only inside functions
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(Logintoken)
	fmt.Printf("Variables is of type : %T \n", Logintoken)

}
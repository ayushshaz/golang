package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["js"]="javascript"
	languages["rb"]="ruby"
	languages["py"]="python"

	fmt.Println("List of all languages", languages)
	fmt.Println("Js shorts for ", languages["js"])

	delete(languages,"rb")
	fmt.Println("List of all languages", languages)


	// loops in golang
	for key, value := range languages {
		fmt.Printf("Key is %v, Value is %v",key, value)
	}
	for _, value := range languages {
		fmt.Printf("Value is %v", value)
	}
}
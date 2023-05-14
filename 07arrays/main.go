package main

import "fmt"

func main() {
	fmt.Println("Welcome to array section")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Mango"
	fruitList[3] = "Guava"
	fmt.Println("Fruit list is: ",fruitList)
	fmt.Println("Fruit list is: ",len(fruitList))

	var vegList  = [3]string{"Potato","mushroom","beans"}
	fmt.Println("Vegetable list is: ",vegList)
}
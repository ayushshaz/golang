package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in Golang")

	var fruitList = []string{"Tomato","Apple","Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango","Papaya")
	fmt.Println("Type of fruitlist is :", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("Type of fruitlist is :", fruitList)

	highScores := make([]int, 4)

	highScores[0] =234
	highScores[1] =235
	highScores[2] =236
	highScores[3] =237

	highScores = append(highScores, 238,239,240)
	fmt.Println("Values of Highscores:", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

// how to remove an index from slice
	var courses = []string{"Ayush","Mayank","Adarsh","Shivang","Saksham","Pandey"}
	fmt.Println("Courses values",courses)
	var index = 3	// Remove Shivang
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println("Courses values",courses)
}

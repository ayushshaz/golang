package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch-case in Golang")

	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1

	switch dice {
	case 1:
		fmt.Println("1 on dice")
	case 2:
		fmt.Println("2 on dice")
	case 3:
		fmt.Println("3 on dice")
		fallthrough // it triggers the next case
	case 4:
		fmt.Println("4 on dice")
		fallthrough
	case 5:
		fmt.Println("5 on dice")
	case 6:
		fmt.Println("6 on dice")
	default:
		fmt.Println("Nothing")
	}

	// Loop
	// break continue works an and even goto-> goto label  label:
}

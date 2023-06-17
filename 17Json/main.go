package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name 	 string `json:"coursename"`// alias
	Price 	 int
	Platform string
	Password string		`json:"-"`
	tags 	 []string	`json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcome to createJson")
	Encodejson()
}

func Encodejson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "learncodeonline.in", "ayu1234", []string{"web-dev","js"}},
		{"Fullstack Bootcamp", 199, "learncodeonline.in", "ayu1234", []string{"full-stack","js"}},
		{"Google Bootcamp",  99, "learncodeonline.in", "ayus1234",nil },
	}

	// Package this data as json
	finalJson, err := json.Marshal(lcoCourses)
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

	// Package this data as json with some changes
	finalJsonIndent, err := json.MarshalIndent(lcoCourses,"","\t")
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJsonIndent)

} 

func DecodeJson(){
	// jsonDataFromWeb := []byte
}
package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=id1234"

func main() {
	fmt.Println("Welcome to URLS in Golang")
	fmt.Println(myurl)

	// parsing, check it url is correct
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
    // the Query part in result is stored like Map.
	fmt.Println(qparams["coursename"])

	for _, val:= range qparams {
		fmt.Println("Param is", val)
	}

	partsOfURL := &url.URL{   // pass it as reference without reference also it works fine
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
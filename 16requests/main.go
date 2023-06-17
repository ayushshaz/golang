package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
// url = scheme//host:port/path?query
// url = "https://lca.dev:3000/learn?coursename=reactjc&paymentid=123121"
func main() {
	partsOfurl := &url.URL{ // url should be used in reference
		Scheme: "https",
		Host: "lca.dev",
		Path: "/tutcuss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfurl.String()
	fmt.Println(anotherURL)

	fmt.Println("Welcome to Get request")
	// PerformgetRequest()
	// PerformPostRequest()
	PerformPostformRequest()
}

func PerformgetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
// First way
	fmt.Println(string(content))  
// Second way
	var responseString strings.Builder // responsestring holds the raw data and its method can be used directly
	byteCount, _ := responseString.Write(content)
	fmt.Println("Bytecount is ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"
	// json payload
	requestbody := strings.NewReader(`
		{
			"Courname":"Golang",
			"Name":"Ayush"
		} 
	`)
	response, err := http.Post(myurl, "application/json", requestbody) // application/json is being as info that its json
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostformRequest() {
	const myurl = "http://localhost:8000/postform"
	// Form data
	data := url.Values{}
	data.Add("firstname", "Ayush")
	data.Add("lastname", "Sharma")
	data.Add("email", "ayush@gmail.com")

	response,err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web request")
	response, err := http.Get(url)	// variable response points to actual response and actual response can be manipulated
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close()
	fmt.Println(response)
	// &{200 OK 200 HTTP/2.0 2 0 map[Accept-Ranges:[bytes] Age:[30947] Cache-Control:[public, max-age=0, must-revalidate] Content-Type:[text/html; charset=UTF-8] Date:[Tue, 23 May 2023 05:39:05 GMT] Etag:["b3141ad8710e6c62813315548dea5d0e-ssl-df"] Server:[Netlify] Strict-Transport-Security:[max-age=31536000] Vary:[Accept-Encoding] X-Nf-Request-Id:[01H14FCYH11PP3RTFCA9YR0YV5]] 0xc0001566f0 -1 [] false true map[] 0xc000138000 0xc000126420}
	fmt.Println("------------------------")
	databytes, err := ioutil.ReadAll(response.Body) 
	if err != nil {
		panic(err)
	}
	fmt.Println(databytes)     // In raw byte type
	fmt.Println("------------------------")
	content := string(databytes)
	fmt.Println(content)
}
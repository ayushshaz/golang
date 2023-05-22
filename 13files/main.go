package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "This Content needs to go in a file"

	file, err := os.Create("./mylcogofile.txt")
	if(err!= nil){
		panic(err)
	}

	length,err := io.WriteString(file,content)
	if(err!= nil){
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}
func readFile(filename string)  {
	databyte, err := ioutil.ReadFile(filename)
	if err!=nil {
		panic(err)
	}
	fmt.Println("text inside data is ", databyte) // this will give in bytes
	fmt.Println("text inside data is ", string(databyte))// will convert in string and show
}
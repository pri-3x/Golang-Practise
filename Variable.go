package main

import (
	"fmt"
)

var i int = 27

func main() {

	fmt.Println(i)

	var i int = 42
	var theHTTP string = "www.google.com"

	fmt.Println(i)
	fmt.Println(theHTTP)
	fmt.Printf("%T", i)
	fmt.Printf("%T", theHTTP)
}

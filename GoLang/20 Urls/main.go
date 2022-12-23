package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&price=230"

func main() {
	fmt.Println("Working with Urls")
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
}

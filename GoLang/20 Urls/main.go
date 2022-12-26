package main

import (
	"fmt"
	"net/url"
)

const Myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=gbhvalue120"

func main() {
	fmt.Println("URL's In Golang")
	fmt.Println(Myurl)
	fmt.Printf("The type of Url is : %T", Myurl)

	// parsing
	content, _ := url.Parse(Myurl)
	fmt.Println(content.Scheme)
	fmt.Println(content.Host)
	fmt.Println(content.Path)
	fmt.Println(content.RawQuery)
	fmt.Println(content.Port())

	// the RawQuery data is stored in form of map and we can acces it via
	data := content.Query()
	fmt.Println(data)
	// fmt.Printf("print %T", data["coursename"])

	for _, val := range data {
		fmt.Println("Param is ", val)
	}

	urlInfo := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=supraja",
	}
	anotherurl := urlInfo.String()
	fmt.Println(anotherurl)

	// now we trying to create a url based info

}

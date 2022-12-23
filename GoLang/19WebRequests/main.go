package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://github.com/Manda-supraja26"

func main() {
	fmt.Println("Web Requests in Golang")
	response, err := http.Get(url)
	CheckError(err)
	databytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	content := string(databytes)
	fmt.Println(content)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}

}

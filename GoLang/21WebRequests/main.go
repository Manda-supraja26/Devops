package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("WebRequests in Golang")
	PerformGetRequest()

}
func PerformGetRequest() {
	const url = "http://localhost:4000/"

	response, err := http.Get(url)
	CheckError(err)
	defer response.Body.Close()

	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length")

	content, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	//  using string builder

	var responseBody strings.Builder

	body, err := responseBody.Write(content)
	CheckError(err)
	fmt.Println("body:", body)
	fmt.Println("content", responseBody.String())

	// fmt.Println(string(content))

}

func PerformPostRequest() {
	const url = "http://localhost:4000/"

	body := strings.NewReader(`
	{
		"name": "supraja",
		"email": "mandasupraja1365139@gmail.com"
	}
	`)

	response, err := http.Post(url, "application/json", body)
	CheckError(err)

	defer response.Body.Close()
	fmt.Println(response)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

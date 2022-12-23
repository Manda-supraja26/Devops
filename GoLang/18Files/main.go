package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Golang")

	content := "hey we are working with files"
	file, err := os.Create("./text.txt")
	if err != nil {
		panic(err)
	}

	// io.WriteString(file,content)

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", length)
	defer file.Close()

	readFile("./text.txt")

}

func readFile(filname string) {
	databytes, err := ioutil.ReadFile(filname)
	CheckError(err)
	fmt.Println(string(databytes))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

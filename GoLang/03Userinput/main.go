package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our code: ")

	// what ever the reader variable is reading i want to store in a variable
	// Here where we are using Comma ok // err ok

	input, err := reader.ReadString('\n')
	fmt.Println("thank you for rating:", input)
	fmt.Printf("thank you for rating: %T", input)
	fmt.Println("error is :", err)
}

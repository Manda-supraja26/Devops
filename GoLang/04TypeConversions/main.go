package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome To our Organization!")

	// Defined the reader to read from the user with the help of bufio and os
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("your Rating pls :")
	// we are storing the data read by the reader into a input variable
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating", input)
	fmt.Printf("type of input is %T", input)

	// numRating := input + 1 this won't work because input is string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numRating + 1)

	}

}

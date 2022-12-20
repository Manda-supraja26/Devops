package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch cases in Golang")
	rand.Seed(time.Now().Unix())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is :", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can move two 1 spot")
	case 2:
		fmt.Println("you can move two 2 spots")
	case 3:
		fmt.Println("you can move two 3 spots")
		fallthrough
	case 4:
		fmt.Println("you can move two 4 spots")
	case 5:
		fmt.Println("you can move two 5 spots")
	case 6:
		fmt.Println("you can move two 6 spots")
	default:
		fmt.Println("hey you are out our dice value")

	}
}

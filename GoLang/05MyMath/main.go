package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// random Number
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(5))

	// fmt.Print(time.Now())

	// random from crypto

	myRandom, _ := rand.Int(rand.Reader, big.NewInt(50))
	fmt.Println("the random number is :", myRandom)
}

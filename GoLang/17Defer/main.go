package main

import "fmt"

func main() {
	defer fmt.Println("Defer")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Class")
	// when we use defer key word then that is compiled at the last
	defer Numbers()

}

func Numbers() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}

}

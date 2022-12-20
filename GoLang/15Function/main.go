package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	// Add()

	// func Great()  {
	// 	fmt.Println("Hey! This is Nested Function")

	// }
	Great()
	// we can't do this but instead we can define it out side and call it in main
	result := Add(9, 10)
	fmt.Println("Result is:", result)

	ans := AddAll(1, 2, 3, 4, 5, 6)
	fmt.Println(ans)
}

func Great() {
	fmt.Println("Hey! This is Nested Function")

}

func Add(val1 int, val2 int) int {
	return val1 + val2
}

func AddAll(val ...int) int {
	total := 0
	for _, v := range val {
		total += v
	}
	return total
}

package main

import "fmt"

func main() {
	fmt.Println("welcome to class on Arrays in Golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "grapes"
	fruitList[3] = "watermelon"

	fmt.Println("fruit list is : ", fruitList)
	fmt.Println("length of array", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(vegList)

}

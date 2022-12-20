package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on Slice")

	var fruitList = []string{"apple"}
	fmt.Printf("Type of fruitList is %T ", fruitList)
	fmt.Println(len(fruitList))

	fruitList = append(fruitList, "mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	marks := make([]int, 4)
	marks[0] = 90
	marks[1] = 99
	marks[2] = 98
	marks[3] = 96
	// marks[4] = 87 we cant do assign index which is out of bound
	// but can append the value into the marks
	marks = append(marks, 90, 95, 35)
	// [90 99 98 96 90 95 35]

	fmt.Println(marks)

	// sort method with slice
	sort.Ints(marks)
	fmt.Println(marks)
	// [35 90 90 95 96 98 99]

	// how to delete a value from slices based on index

	var names = []string{"Tom", "bob", "jhon", "Emma"}
	fmt.Println(names)

	var index int = 2
	names = append(names[:index], names[index+1:]...)
	fmt.Println(names)

}

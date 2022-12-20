package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thusrday", "Friday", "Saturday"}
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// 	if days[i] == "Monday" {
	// 		var Num string = "day"
	// 		days = append(days, Num)
	// 	}

	// }

	// break
	// Continue

	for i := range days {
		fmt.Println(days[i])
		if i == 6 {
			goto supraja
		}
	}

	// for index, day := range days {
	// 	fmt.Print(index + 1)
	// 	fmt.Println(day)
	// }
supraja:
	fmt.Println("Hey! This is Supraja's day")

	value := 1
	for value <= 10 {
		fmt.Println(value)
		value++
	}

	// Go to

}

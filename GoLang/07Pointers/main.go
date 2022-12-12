package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on pointers")

	// why pointers?
	// some time when we are passing the reference of the variable copy of variable is passed and there might be some casese where you fail to pass copy of reference variable
	// in that cases we can pass pointers ,pointers stores memory address so it guarenty that  actual value is passed

	// var pntr *int
	// fmt.Println(pntr) --> value is nil

	number := 23
	var pntr = &number
	fmt.Println("value of actual pointer is ", pntr)
	fmt.Println("value of actual pointer is ", *pntr)

	// Welcome to class on pointers
	// value of actual pointer is  0xc000016088
	// value of actual pointer is  23

	*pntr = *pntr * 2
	fmt.Println(*pntr)
	fmt.Println(number)

}

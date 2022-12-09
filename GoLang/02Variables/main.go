package main

import (
	"fmt"
)

// number := 345  is not acceptable outside the function only
// works in functions or methods only

// to intilize varible outside function we need to use
var number int = 90

// declaring constants
const LoginToken string = "dhjehj"

// variables starting with Capital letter is a public
func main() {
	var username string = "supraja"
	fmt.Print(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = false
	// fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal int = 257
	fmt.Println(smallVal)
	fmt.Printf("variable is of type : %T \n", smallVal)

	var smallFloatVal float64 = 257.099839819787535676
	fmt.Println(smallFloatVal)
	fmt.Printf("variable is of type : %T \n", smallFloatVal)

	//  default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type : %T", anotherVariable)

	var Var2 string
	// fmt.Printf("this is :", Var2)
	// fmt.Printf("supraja+ %T", Var2)
	fmt.Printf("variable is of type : %T", Var2)
	fmt.Println()

	// implicit way to declare variables
	var website = "supraja"
	fmt.Println(website)
	website = "lala"
	fmt.Println(website)

	// No var style
	numberofusers := 2000.9
	fmt.Println(numberofusers)
}

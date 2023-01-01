package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var ConferenceName string = "Go Conference"
	const ConferenceTicket = 50
	var RemainingTickets uint = 50

	fmt.Println("Welcome to ", ConferenceName, " booking application")
	fmt.Println("Get Your ticket here to attend")
	fmt.Printf("Total Tickets %v and %v  tickets are available to Book\n", ConferenceTicket, RemainingTickets)

	// Taking username and number of tickets from the user using scan
	// we can also use buffio
	// var firstName string
	// var lastName string
	// var email string
	// var userTickets uint

	// fmt.Println("Enter Your FirstName :")
	// fmt.Scan(&firstName)

	// fmt.Println("Enter Your LastName :")
	// fmt.Scan(&lastName)

	// fmt.Println("Enter Your Email :")
	// fmt.Scan(&email)

	// fmt.Println("Enter number of tickets :")
	// fmt.Scan(&userTickets)

	// RemainingTickets = RemainingTickets - userTickets

	// fmt.Printf("Thank you %v %v for booking %v tickets. You will recive the conformation email at %v. \n", firstName, lastName, userTickets, email)
	// fmt.Printf("Remaining tickests available are %v", RemainingTickets)

	// reader := bufio.NewReader(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our code: ")

	// what ever the reader variable is reading i want to store in a variable
	// Here where we are using Comma ok // err ok

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(input)

}

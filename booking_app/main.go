package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here !")
	fmt.Printf("Total tickets: %d \n Remaining Tickets: %d \n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("%s %s (%s) booked %d tickets \n", firstName, lastName, email, userTickets)
	

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here !")
	fmt.Printf("Total tickets: %d \n Remaining Tickets: %d \n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int

	for {
		fmt.Print("Enter first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("array: %v \n", bookings)
		fmt.Printf("%s %s (%s) booked %d tickets \n", firstName, lastName, email, userTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])

		}

		if remainingTickets == 0 {
			fmt.Printf("Fully booked \n")
			break

		}

		fmt.Printf("first name list: %v \n", firstNames)

	}

}

package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var remainingTickets = 50

func greetUsers() {

	var conferenceName string = "Go Conference"

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here !")
	fmt.Printf("Total tickets: %d \n Remaining Tickets: %d \n", conferenceTickets, remainingTickets)

}

func userDetails() {

	var firstName string
	var lastName string
	var email string
	var userTickets int
	var bookings []string

	for {

		fmt.Print("Enter first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter last name: ")
		fmt.Scan(&lastName)

		fmt.Print("Enter email: ")
		fmt.Scan(&email)

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets

		if isValidEmail && isValidName && isValidUserTicket {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
			fmt.Printf("%s %s (%s) booked %d tickets \n", firstName, lastName, email, userTickets)
			fmt.Printf("Tickets remaining %d \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])

			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name is invalid")
			}
			if !isValidEmail {
				fmt.Println("Please enter a valid email")

			}
			if !isValidUserTicket {
				fmt.Println("Please enter a valid ticket number")
			}
		}

		if remainingTickets == 0 {
			fmt.Printf("Fully booked \n")
			break
		}

	}
}

func main() {

	greetUsers()
	userDetails()
}

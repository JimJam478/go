package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var remainingTickets = 50

var bookings = make([]UserDetails, 0)

type UserDetails struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func greetUsers() {

	var conferenceName string = "Go Conference"

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("Get your tickets here !")
	fmt.Printf("Total tickets: %d \n Remaining Tickets: %d \n", conferenceTickets, remainingTickets)

}

func getUserDetails() (string, string, string) {

	var firstName string
	var lastName string
	var email string

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter email: ")
	fmt.Scan(&email)

	return firstName, lastName, email
}

func bookTickets() {

	var userTickets int

	for {
		firstName, lastName, email := getUserDetails()
		fmt.Print("Enter number of tickets:")
		fmt.Scan(&userTickets)

		var userData = UserDetails{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: userTickets,
		}

		bookings = append(bookings, userData)

		isValidName, isValidEmail, isValidUserTicket := userValidation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidUserTicket {
			remainingTickets = remainingTickets - userTickets
			fmt.Printf("%s %s (%s) booked %d tickets \n", firstName, lastName, email, userTickets)
			fmt.Printf("Tickets remaining %d \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking.firstName)
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
	bookTickets()

}

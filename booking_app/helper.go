package main

import "strings"

func userValidation(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTicket := userTickets > 0 && userTickets <= remainingTickets
	return isValidEmail, isValidName, isValidUserTicket
}

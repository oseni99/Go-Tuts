package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

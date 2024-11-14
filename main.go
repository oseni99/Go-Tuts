package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var emailAddress string
		var userTickets uint
		// ask the user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&emailAddress)
		fmt.Println("Enter number of Tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 3 && len(lastName) >= 3
		isValidEmail := strings.Contains(emailAddress, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, emailAddress)
			fmt.Printf("%v tickets remaining for the %v are left\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) // split the names in that array/slice
				// var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v", firstNames)

			if remainingTickets == 0 {
				// end the program
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Kindly check your first name and last name for changes. Both must be at least 3 characters long.")
			} else if !isValidEmail {
				fmt.Println("Kindly enter a valid email address that contains an '@' symbol.")
			} else if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid. Ensure you are booking at least 1 ticket and not more than the available tickets.")
			}
		}

	}

}

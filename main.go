package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// := does not work for global variables in Go i.e packages
var conferenceName = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// struct will give you a predefined structure
//
//	use struct if you know the values in the dicts but map if unknown
type UserData struct {
	firstName       string
	lastName        string
	emailAddress    string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, emailAddress, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, emailAddress)
		//  make it concurrent i.e multi threaded with just go
		//  The default go routine doesnt tell it to wait so we have to teo=ll it itslf how many go routines it should wait for
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAddress)

		//  call func to print first names
		firstNames := FirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end the program
			fmt.Println("Our conference is booked out. Come back next year")
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
	//  basically tells the code to wait until all routines are done before closing
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// In Go i have to specify the exact type of what im returning when i use return statement
func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

//  You can return multiple functions in go which is super amazing

func getUserInput() (string, string, string, uint) {
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
	return firstName, lastName, emailAddress, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	//  creating a map i.e dictionary
	//  map creation is done with that make
	// maps in go is limited to only one data type

	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		emailAddress:    emailAddress,
		numberOfTickets: userTickets,
	}
	// Map example
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["emailAddress"] = emailAddress
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for the %v are left\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, FirstName string, lastName string, emailAddress string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, FirstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, emailAddress)
	fmt.Println("###############")
	//  This now removes the thread that was started after it finished executing
	wg.Done()
}

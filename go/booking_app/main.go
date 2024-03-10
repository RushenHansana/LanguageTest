package main

import (
	"fmt" //format package

	

	//importing the helper package
	"booking-app/helper"
)

// package level variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //initialize an empty slice of map

//create a custom data type
type UserData struct {
	firstName string
	lastName  string
	email     string
	numberOfTickets   uint

}

// which line executes starts
func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {

			bookTickets(userTickets, firstName, lastName, email)

			//call the function print firstnames
			FirstNames := getFirstNames()
			fmt.Printf("The first names of all the bookings are %v \n", FirstNames)

			if remainingTickets == 0 {
				fmt.Println("We are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("first name and last name you entered is too short, please try again \n")
			}
			if !isValidEmail {
				fmt.Printf("Email you entered don't hava @ sign, please try again \n")
			}
			if !isValidTicketsNumber {
				fmt.Printf("Invalid number of tickets, please try again \n")
			}

		}

	}
	city := "New York"

	switch city {
	case "New York", "NYC":
		fmt.Println("Welcome to New York")
	case "London":
		fmt.Println("Welcome to London")
	default:
		fmt.Println("Where are you?")
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now!")

}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Please enter your first name")
	fmt.Scanln(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scanln(&lastName)
	fmt.Println("Please enter your email")
	fmt.Scanln(&email)
	fmt.Println("Please enter the number of tickets you want to book")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//create a map for a user we cannot use mixed data types in a map
	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		numberOfTickets:   userTickets,
	}


	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v \n" ,bookings)

	fmt.Printf("Thank you %v %v, you have successfully booked %v tickets for the %v conference.You will recieve confirmation at %v \n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("We have %v tickets remaining \n", remainingTickets)
}

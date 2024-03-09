package main

import (
	"fmt" //format package
	"strings"
)

// which line executes starts
func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now!")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name")
		fmt.Scanln(&firstName)

		fmt.Println("Enter your last name")
		fmt.Scanln(&lastName)

		fmt.Println("Enter your email")
		fmt.Scanln(&email)

		fmt.Println("How many tickets do you want to book?")
		fmt.Scanln(&userTickets)

		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)
	
			fmt.Printf("Thank you %v %v, you have successfully booked %v tickets for the %v conference.You will recieve confirmation at %v \n", firstName, lastName, userTickets, conferenceName, email)
			fmt.Printf("We have %v tickets remaining \n", remainingTickets)
	
			firstNames := []string{}
	
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				fmt.Printf("The first name of the booking is %v \n", names[0])
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("The first names of all the bookings are %v \n", firstNames)
	
			noTicketsRemaining := remainingTickets == 0
	
			if noTicketsRemaining {
				fmt.Println("We are sold out!")
				break
			}
		} else {
			fmt.Printf("Sorry, we only have %v tickets remaining \n", remainingTickets)
			
		}


	}

}

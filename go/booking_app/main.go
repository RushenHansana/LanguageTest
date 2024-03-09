package main

import "fmt" //format package

//which line executes starts
func main() {

	//variable declaration if we don't use it, it will give error
	var conferenceName = "Go Conference"
	//alternative way to declare variable
	conferenceDate := "2021-12-01"

	//variable declaration with value that can't be changed
	const conferenceTickets int = 50  
	//setting the type of the variable ensures that the variable can only store the value of that type
	var remainingTickets int = 50
	//only one main function is allowed

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	//space gets automatically added between the strings and variables
	fmt.Printf("Welcome to %v booking application\n", conferenceName) 
	//%v is a placeholder for the variable
	fmt.Printf("We have total of %v tickets and %v tickets are remaining.\n", conferenceTickets, remainingTickets) 
	//fmt.Printf is used to format the string	
	fmt.Printf("Get your tickets here to attend on %v\n", conferenceDate)	

	//array declaration with size of 50
	// var bookings = [50]string{}

	var bookings [50]string

	//slices vs array dynamic sized array which can grow or shrink
	var books []string
	books = append(books, "hello")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user to enter the name
	//& is used to get the address of the variable pointer
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)//assign the value to the variable memory location

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book:")
	fmt.Scan(&userTickets)

	// fmt.Print(remainingTickets)
	// fmt.Print(&remainingTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("Whole array is: %v \n", bookings)
	fmt.Printf("first element of array is: %v \n", bookings[0])
	fmt.Printf("type elements of array is: %T \n", bookings)
	fmt.Printf("length of array is: %v \n", len(bookings))
	
	fmt.Printf("User %v has booked %v tickets & you will reciewve email at %v  \n", firstName, userTickets, email)
	fmt.Printf("Remaining tickets are %v\n", remainingTickets)

	fmt.Printf("Wholeslice is: %v \n", books)

	
	

}



package main

import (
	"fmt"
)

func main() {
	var conferanceName = "Go Conferance"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %s booking application\n", conferanceName)
	fmt.Println("we have total of", conferenceTickets, "tickets and", remainingTickets, "are still avilable")
	fmt.Println("Book your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var bookings []string

		fmt.Println("Enter your First Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastname")
		fmt.Scan(&lastName)

		fmt.Println("Enter the tickets you want to book:")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thanks %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferanceName)

		fmt.Printf("Type of bookins is %T\n", bookings)
		//firstNames := []string{}
		//for index, booking := range bookings {
		//	var names = strings.Fields(booking)
		//	var firstName := names[0]
	}
	fmt.Print("Here is slices")
}

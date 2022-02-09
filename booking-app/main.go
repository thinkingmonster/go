package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings []UserData

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidUserTickets := helper.ValidateUserInput(firstName,
			lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)
			initials := getFirstNames()
			fmt.Printf("These are our bookings %v\n", initials)
			if remainingTickets == 0 {
				//end program
				fmt.Println("Our conference is booked out.Come back next year")
				break
			}
		} else {
			switch false {
			case isValidName:
				fmt.Println("You have provided wrong first or last name")
			case isValidEmail:
				fmt.Println("You have provided invalid email")
			case isValidUserTickets:
				fmt.Println("Wrong number of tickets provided")

			}
		}
	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	var initials []string
	for _, data := range bookings {
		initials = append(initials, data.firstName)
	}
	return initials
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("\n")
	// ask user for their name
	fmt.Printf("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your Email address: ")
	fmt.Scan(&email)

	fmt.Printf("How many tickets you want to book: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int,
	firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("Thank you %v %v for booking %v Tickets,You will recieve confirmation on your email: %v\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("Total Tickets remaining for %v: %v\n", conferenceName, remainingTickets)
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	// Simulating delay in this function
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("#######################")
	wg.Done()
}

package helper

import (
	"strings"
)

func ValidateUserInput(firstName string,
	lastName string,
	email string,
	userTickets int,
	remainingTickets int,

) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidUserTickets := userTickets <= remainingTickets && userTickets > 0
	return isValidName, isValidEmail, isValidUserTickets

}

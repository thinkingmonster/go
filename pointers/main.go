package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to ", myString)
	changeUsingPointer(&myString)
	log.Println("myString after calling function, set to ", myString)

}

func changeUsingPointer(s *string) {
	*s = "Red"
}

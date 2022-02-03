package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	UserName        string
	Age             int
	FavouriteNumber float64
	hasDog          bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("What is your name?")
	user.Age = readInt("What is your age?")
	user.FavouriteNumber = readFloat("What is your Favourite number")
	user.hasDog = haveDog("Do have a dog?")
	fmt.Printf("Your name is %s. Your are %d\nIt looks like that your fav number is %.2f and you have a dog: %t\n",
		user.UserName,
		user.Age,
		user.FavouriteNumber,
		user.hasDog)

}

func takeInput() bool {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false

		} else {
			fmt.Println("Please enter Yes or no")
		}
	}
}

func prompt() {
	fmt.Print("--> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "" {
			fmt.Println("Please input your name")

		} else {
			return userInput
		}
	}

}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Print("Please enter a whole number")
		} else {
			return num
		}

	}

}

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)
		num, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Print("Please enter a whole number")
		} else {
			return num
		}

	}

}

func haveDog(s string) bool {
	fmt.Println(s)
	prompt()
	return takeInput()

}

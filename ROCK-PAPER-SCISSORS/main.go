package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	rand.Seed(time.Now().UnixNano())
	playerChoice := ""
	playerValue := -1
	playerScore := 0
	computerScore := 0

	reader := bufio.NewReader(os.Stdin)

	clearScreen()
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins.Good luck!")
	fmt.Println()
	for i := 1; i <= 3; i++ {
		computerValue := rand.Intn(3)

		fmt.Printf("------Round %d------\n", i)
		fmt.Print("Please enter rock, paper, or scissors ->")
		playerChoice, _ = reader.ReadString('\n')
		playerChoice = strings.Replace(playerChoice, "\n", "", -1)

		switch playerChoice {
		case "rock":
			playerValue = ROCK
			break
		case "paper":
			playerValue = PAPER
			break
		case "scissors":
			playerValue = SCISSORS
		}

		switch computerValue {
		case ROCK:
			fmt.Println("Computer selected Rock")
			break
		case PAPER:
			fmt.Println("Computer selected Paper")
			break
		case SCISSORS:
			fmt.Println("Computer selected Scissors")

		}

		//select Winner
		if playerValue == computerValue {
			fmt.Println("Its a draw")
		} else {
			switch playerValue {
			case ROCK:
				if computerValue == PAPER {
					fmt.Println("Computer wins")
					computerScore = computerScore + 1
				} else {
					fmt.Println("Player wins")
					playerScore = playerScore + 1
				}
				break
			case PAPER:
				if computerValue == SCISSORS {
					fmt.Println("Computer wins")
					computerScore = computerScore + 1
				} else {
					fmt.Println("Player wins")
					playerScore = playerScore + 1
				}
				break
			case SCISSORS:
				if computerValue == ROCK {
					fmt.Println("Computer wins")
					computerScore = computerScore + 1
				} else {
					fmt.Println("Player wins")
					playerScore = playerScore + 1
				}
				break
			default:
				fmt.Println("Wrong choice")
				i = i - 1

			}
		}
	}
	fmt.Println()
	fmt.Println("-- Final Score --")
	if playerScore > computerScore {
		fmt.Println("Player wins")
	} else if computerScore > playerScore {
		fmt.Println("Computer wins")
	} else {
		fmt.Println("Its a draw")
	}
	fmt.Printf("Player:(%v/3)  Computer:(%v/3)", playerScore, computerScore)

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

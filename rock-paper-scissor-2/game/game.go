package game

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

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// use select to process input in channels
	// Print to screen
	// Keep track of round number.
	for {

		select {
		case round := <-g.RoundChan:
			g.Round.RoundNumber = g.Round.RoundNumber + round

		case msg := <-g.DisplayChan:
			fmt.Println(msg)
		}

	}
}

// ClearScreen clears the screen
func (g *Game) ClearScreen() {
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

func (g Game) PlayerIntro() {
	fmt.Println("Rock, Paper & Scissors")
	fmt.Println("-----------------------")
	fmt.Println("Game is played for three rounds, and best of three wins.Good luck!")
	fmt.Println()

}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1
	fmt.Println()
	fmt.Println("Round", g.Round.RoundNumber)
	fmt.Println("--------------")
	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	switch playerChoice {
	case "rock":
		playerValue = ROCK
		break
	case "paper":
		playerValue = PAPER
		break
	case "scissors":
		playerValue = SCISSORS
		break
	default:
		playerValue = -1
	}

	fmt.Println()
	g.DisplayChan <- fmt.Sprintf("Player choose %v", strings.ToUpper(playerChoice))

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

	fmt.Println()

	//select Winner
	if playerValue == computerValue {
		g.DisplayChan <- "Its a draw"
		return false
	} else {
		switch playerValue {
		case ROCK:
			if computerValue == PAPER {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case PAPER:
			if computerValue == SCISSORS {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		case SCISSORS:
			if computerValue == ROCK {
				g.computerWins()
			} else {
				g.playerWins()
			}
			break
		default:
			g.DisplayChan <- "Invalid Choice!"
			return false

		}
	}
	return true
}

func (g *Game) computerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer Wins"
}

func (g *Game) playerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player Wins"
}

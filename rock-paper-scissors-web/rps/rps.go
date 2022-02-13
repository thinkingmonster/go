package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK         = 0
	PAPER        = 1
	SCISSORS     = 2
	PLAYERWINS   = 1
	COMPUTERWINS = 2
	DRAW         = 3
)

type Round struct {
	Winner         string `json:"winner"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult    string `json:"round_result"`
}

func PlayRound(playerValue int) Round {

	winner_msg := map[int]string{
		1: "You are awesome!",
		2: "Champions never loose.",
		3: "You should try a lottry",
	}

	draw_msg := map[int]string{
		1: "oh its a draw!",
		2: "Same mind thinks alike",
		3: "You are computer are so same.",
	}

	loose_msg := map[int]string{
		1: "oopse, you loose",
		2: "Never run for a mayer",
		3: "you are doomed ",
	}
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	messageValue := rand.Intn(3) + 1
	computerChoice := ""
	roundResult := ""
	winner := ""

	switch computerValue {
	case ROCK:
		computerChoice = "Computer selected Rock"
		break
	case PAPER:
		computerChoice = "Computer selected Paper"
		break
	case SCISSORS:
		computerChoice = "Computer selected Scissors"

	}

	if playerValue == computerValue {
		roundResult = "Its a draw"
		winner = draw_msg[messageValue]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		winner = winner_msg[messageValue]
	} else {
		roundResult = "Computer wins!"
		winner = loose_msg[messageValue]
	}
	var result Round
	result.Winner = winner
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult

	return result

}

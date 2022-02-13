package main

import (
	"myapp/game"
)

func main() {
	displayChan := make(chan string)
	roundChan := make(chan int)

	game := game.Game{
		DisplayChan: displayChan,
		RoundChan:   roundChan,
		Round: game.Round{
			RoundNumber:   0,
			PlayerScore:   0,
			ComputerScore: 0,
		},
	}
	go game.Rounds()
	game.ClearScreen()
	game.PlayerIntro()

	for {

		game.RoundChan <- 1
		<-game.RoundChan

		if game.Round.RoundNumber > 3 {
			break
		}
		if !game.PlayRound() {
			game.RoundChan <- -1
			<-game.RoundChan
		}
	}
	//fmt.Println()
	//fmt.Println("-- Final Score --")
	//if playerScore > computerScore {
	//	fmt.Println("Player wins")
	//} else if computerScore > playerScore {
	//	fmt.Println("Computer wins")
	//} else {
	//	fmt.Println("Its a draw")
	//}
	//fmt.Printf("Player:(%v/3)  Computer:(%v/3)", playerScore, computerScore)

}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press enter when ready"

func main() {

	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)
}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Guess a number game")
	fmt.Println("----------------------------")
	fmt.Println("")

	fmt.Println("Think  of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply the result  by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you orignally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", substraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)

}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"myapp/mylogger"
)

func main() {
	//likeWhile()
	//InfiniteLoop()
	//NestedLoop()

	for i := 0; i != -1; i++ {
		fmt.Println("I am infinite")
	}
}

func likeWhile() {
	rand.Seed(time.Now().UnixNano())
	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println("i is", i)
		if i > 100 {
			fmt.Println("I is", i, "so loop keeps going")
		}

	}
	fmt.Println("Got", i, "and break out of loop")

}

// InfiniteLoop Reads the input from user 5 times and prints the same.
// I have created a ch , a logger that runs  forever listening for the input.
func InfiniteLoop() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go mylogger.ListenForLog(ch)
	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second) // had to put sleep as ch was printing too quick even before the prompt.

	}

}

func NestedLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Print("Value of i is ", i)
		for j := 0; j < 3; j++ {
			fmt.Println(" and j is ", j)
		}

	}

}

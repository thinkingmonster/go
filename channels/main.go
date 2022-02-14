package main

import (
	"log"
	"myapp/helper"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helper.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	num := <-intChan
	log.Println(num)
}

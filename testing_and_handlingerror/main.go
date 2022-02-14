package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	result, err := DivideNumbers(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(result)

}

func DivideNumbers(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("can not divide by zero")
	}
	result := x / y
	return result, nil
}

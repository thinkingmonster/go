package main

import "fmt"

func main() {
	var x int = 1
	var y int = 2
	y, x = x, y

	fmt.Println(x, y)

}

package main

import (
	"myapp/packageone"
)

var myVar = "i am mainpackage var"

func main() {
	var blockVar = "I am a block var"
	packageone.PrintMe(blockVar, myVar)

}

package main

import "fmt"

// Working with arrays

func main() {
	//arrays()
	slicesLearn()
}

func arrays() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	fmt.Println(myArray)

	var anotherArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(anotherArray)

	var array1 = [...]int{1, 2, 3, 4}
	var array2 = [4]int{1, 2, 3, 4}

	fmt.Println(array1 == array2)

	array3 := [3]int{1, 2, 3}
	fmt.Println(array3)
}

func slicesLearn() {
	var x []int
	x = append(x, 1, 2, 3)
	fmt.Println(x)

	var x1 = []int{1, 2, 3}
	fmt.Println(x1)

	x2 := []int{4, 5, 6}
	fmt.Println(x2)

}

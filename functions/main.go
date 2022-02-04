package main

import "fmt"

func main() {
	fmt.Println(addMany(1, 2, 3, 4, 5, 6))
	fmt.Println(addMany(1, 2, 3))

	var dog Animals
	dog.Name = "dog"
	dog.Sound = "woof"
	dog.NumberOfLegs = 4
	dog.says()

	cat := Animals{
		Name:         "cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}
	cat.says()

}

type Animals struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

func (a Animals) says() {
	fmt.Printf("The animal %s produces %s sound and has %d legs\n", a.Name, a.Sound, a.NumberOfLegs)
}

// variadic functions
func addMany(num ...int) int {
	total := 0
	for _, i := range num {
		total = total + i
	}
	return total
}

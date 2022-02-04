package main

import "fmt"

type Animals interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	Name     string
	Sound    string
	NoOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NoOfLegs
}

type Cat struct {
	Name     string
	Sound    string
	NoOfLegs int
	HasFur   bool
}

func (c *Cat) Says() string {
	return c.Sound
}

func (c *Cat) HowManyLegs() int {
	return c.NoOfLegs
}

func main() {
	var d Dog
	d.Name = "Jimmy"
	d.Sound = "woof-woof"
	d.NoOfLegs = 4
	riddle1 := Riddle(&d)
	fmt.Printf(riddle1)

	var c Cat
	c.Name = "Tommy"
	c.Sound = "meow-meow"
	c.NoOfLegs = 4
	riddle2 := Riddle(&c)
	fmt.Printf(riddle2)
}
func Riddle(a Animals) string {
	return fmt.Sprintf("I say %s and i have %d legs,who am i\n", a.Says(), a.HowManyLegs())
}

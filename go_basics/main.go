package main

import "fmt"

//aggregate types(array, struct)
// Referance type ( pointers, slices,maps,functions,channels)
// interface types

//Struct

//type Car struct {
//	NumberOfTires int
//	Luxury        bool
//	BucketSeats   bool
//	Make          string
//	Model         string
//	Year          int
//}

func main() {
	//var myStrings [3]string
	//myStrings[0] = "cat"
	//myStrings[1] = "dog"
	//myStrings[2] = "fish"
	//
	//fmt.Println(myStrings)
	//
	////var myCar Car
	////myCar.Make = "Maruti"
	//
	//myCar := Car{
	//	NumberOfTires: 4,
	//	Luxury:        true,
	//}
	//
	//fmt.Println(myCar.NumberOfTires)

	////Pointers
	//x := 10
	//locationOfx := &x
	//
	//fmt.Println(x)
	//fmt.Println(locationOfx)
	//
	//*locationOfx = 20
	//fmt.Printf("The value of x is now %d", x)

	////Slices
	//var animals []string
	//animals = append(animals, "dog")
	//animals = append(animals, "fish")
	//animals = append(animals, "cat")
	//animals = append(animals, "horse")
	////fmt.Println(len(animals))
	////fmt.Println(animals)
	////fmt.Println(animals[:len(animals)-2])
	//animals = DeleteFromSlice(animals, 2)
	//fmt.Println(animals)

	// Maps

	var myMap = make(map[string]int)
	myMap["One"] = 1
	myMap["Two"] = 2
	myMap["Three"] = 3
	myMap["Four"] = 4
	myMap["Five"] = 5

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	for key, value := range myMap {
		fmt.Println(key, value)
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	delete(myMap, "Five")
	fmt.Println(myMap)

	value, ok := myMap["Four"]
	if ok {
		fmt.Printf("Is in map and value is %d", value)
	} else {
		fmt.Println("It is not in map")
	}

}

//func DeleteFromSlice(s []string, index int) []string {
//
//	s[index] = s[len(s)-1]
//	fmt.Println(s[:len(s)-1])
//	s = s[:len(s)-1]
//	fmt.Println(s)
//	return s
//}

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[{
	"first_name": "clark",
	"last_name": "kent",
	"hair_color": "black",
	"has_dog": true
}, {
	"first_name": "Bruce",
	"last_name": "Wayne",
	"hair_color": "black",
	"has_dog": false
}]
`
	//converting json to struct
	var data []Person
	err := json.Unmarshal([]byte(myJson), &data)
	if err != nil {
		log.Println(err)

	}
	fmt.Println(data)

	//converting struct to json

	var anew []Person

	person1 := Person{
		FirstName: "akash",
		LastName:  "thakur",
		HairColor: "black-brown",
		HasDog:    false,
	}

	person2 := Person{
		FirstName: "anjna",
		LastName:  "thakur",
		HairColor: "black",
		HasDog:    false,
	}

	anew = append(anew, person1, person2)
	output, _ := json.MarshalIndent(anew, "", "  ")
	fmt.Println(string(output))
}

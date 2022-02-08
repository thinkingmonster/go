package jsonexample

import (
	"encoding/json"
	"fmt"
)

func myJson() {
	type Author struct {
		Name        string
		Age         int
		IsDeveloper bool
	}

	type Book struct {
		Title  string `json:"title"`
		Author Author `json:"author"`
	}

	fmt.Println("Converting struct to json")
	author := Author{
		Name:        "Jon Bodner",
		IsDeveloper: true,
		Age:         32,
	}
	book := Book{
		Title:  "Learning Go",
		Author: author,
	}

	fmt.Printf("%+v\n", book)
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println("Error occurred")
	} else {
		fmt.Println(string(byteArray))
	}

}

func RemoteJson() {
	type Info struct {
		Description string `json:"desc"`
	}
	type SensorReading struct {
		Name        string `json:"name"`
		Capacity    int    `json:"capacity"`
		Time        string `json:"time"`
		Information Info   `json:"information"`
	}
	jsonString := `{"name":"Battery Sensor","capacity":40,"time":"2019-01-21T19:07:28Z","information":{ "desc":"A new sensor"}}`

	var reading SensorReading
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)

}

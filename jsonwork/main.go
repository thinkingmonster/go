package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var releaseApplicationFile string
var decommissionApplicationFile string

type releaseApplication []struct {
	Deploy string `json:"deploy"`

	Service string `json:"service"`
	Version struct {
		Helm           string `json:"helm"`
		Service        string `json:"service"`
		ServiceCurrent string `json:"serviceCurrent"`
	} `json:"version"`
}

type decommissionApplication []struct {
	Service string `json:"service"`
	Version struct {
		Service string `json:"service"`
	} `json:"version"`
}

func takeInputs() {
	if len(os.Args) != 3 {
		fmt.Printf("Usages:\n\tsimba <release-application-file> <decommission-application-file>")
		os.Exit(1)
	} else {
		releaseApplicationFile = os.Args[1]
		decommissionApplicationFile = os.Args[2]
	}
}

func openFile(filename string) []byte {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
func main() {
	takeInputs()
	releaseApplicationByte := openFile(releaseApplicationFile)
	var ra releaseApplication
	raError := json.Unmarshal(releaseApplicationByte, &ra)
	if raError != nil {
		fmt.Println(raError)
	}

	decommissionApplicationByte := openFile(decommissionApplicationFile)
	var da decommissionApplication
	daError := json.Unmarshal(decommissionApplicationByte, &da)
	if daError != nil {
		fmt.Println(daError)
	}
	conflictsFound := false
	for i := 0; i < len(da); i++ {
		for j := 0; j < len(ra); j++ {
			if da[i].Service == ra[j].Service &&
				da[i].Version.Service == ra[j].Version.Service {
				fmt.Printf("ALERT:  Decomission service %s found in Release json\n\tConficting version %s\n",
					da[i].Service,
					da[i].Version.Service)
				conflictsFound = true
			}
		}

	}
	if !conflictsFound {
		fmt.Println("Success: No conflicting json found")
	}
}

package packageone

import "fmt"

var PackageVar = "i am Package var"

func PrintMe(blockVar, myVar string) {
	fmt.Println(blockVar, myVar, PackageVar)

}

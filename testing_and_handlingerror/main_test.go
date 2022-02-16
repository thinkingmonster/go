package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0.0, true},
	{"expect-5", 50.0, 5, 0.0, false},
}

func TestDivisor(t *testing.T) {
	for _, item := range tests {
		if !item.isErr {
			_, err := DivideNumbers(item.dividend, item.divisor)
			if err != nil {
				t.Error("Got an error when we should not")
			}
		} else {
			_, err := DivideNumbers(item.dividend, item.divisor)
			if err == nil {
				t.Error("Dis not got error when we should")
			}

		}

	}

}

//func TestDivideNumbers(t *testing.T) {
//	_, err := DivideNumbers(100.0, 10.0)
//	if err != nil {
//		t.Error("Got an error when we should not")
//	}
//
//}
//
//func TestBadDivide(t *testing.T) {
//	_, error := DivideNumbers(100.0, 0)
//	if error == nil {
//		t.Error(" not getting an error when we should have")
//	}
//
//}

//Command to execute
//// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

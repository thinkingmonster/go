package main

import "testing"

func TestDivideNumbers(t *testing.T) {
	_, err := DivideNumbers(100.0, 10.0)
	if err != nil {
		t.Error("Got an error when we should not")
	}

}

func TestBadDivide(t *testing.T) {
	_, error := DivideNumbers(100.0, 0)
	if error == nil {
		t.Error("Dis not get an error when we should have")
	}

}

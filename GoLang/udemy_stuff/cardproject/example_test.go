package main

import "testing"

//This is for the simple card game that i didnt try

//The name of the function seems arbitrary, but its uppercase
func TestExampleFunc(t *testing.T) {
	//t is the Test Handler
	f := functotest()

	if f != "Here should be the expected value" {
		t.Errorf("Expected expeced value, but got %v", f)
	}
}

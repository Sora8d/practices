package main

import (
	"fmt"
)

type sal struct {
	ex1 int
}

var continuity = []sal{
	{ex1: 123}, {ex1: 345}, {ex1: 456}}

func main() {
	fmt.Println()
	fmt.Printf("%T", continuity[1])
	fmt.Println()
}

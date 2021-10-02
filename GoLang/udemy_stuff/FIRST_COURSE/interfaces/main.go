package main

import "fmt"

type bot interface {
	getGreeting() string
}

type test_stru struct {
	name     string
	language bot
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

type spanishBot struct{}
type englishBot struct{}

func (sb spanishBot) getGreeting() string {
	//Different logics
	return "Hola"
}

func (eb englishBot) getGreeting() string {
	//Different logics
	return "Hello"
}

func main() {
	var sb spanishBot
	//	var eb englishBot
	manuel := test_stru{name: "Jhon", language: sb}
	printGreeting(sb)
	//	printGreeting(eb)
	fmt.Println(manuel)
	printGreeting(manuel.language)
}

//A good way to see more complex interfaces is the Response type inside of the net/http module.

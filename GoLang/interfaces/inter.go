package main

import "fmt"

type Tomas struct {
}

type Matias struct {
}

type luckyboy interface {
	Winlottery() string
}

type unluckyboy interface {
	Winlottery() string
}

func (t *Tomas) Winlottery() {
	fmt.Println("YOU WON YES")
}

func (m *Matias) Winlottery() {
	fmt.Println("No you lost bud")
}

func main() {
	var vault interface{}

	i := 123
	v := "Hallo"

	vault = i
	vault = v

	fmt.Println(vault)

	var 
}
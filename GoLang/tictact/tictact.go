package main

import (
	"fmt"
)

type board struct {
	player1 int
	player2 int
	board   [3][3]int
	over    bool
	winner  string
	turn    int
}

func transformValuesToStrings(arr [3]int) [3]string {
	var formatted_string [3]string
	for index, element := range arr {
		var transformed_value string
		switch element {
		case 2:
			transformed_value = "O"
		case 1:
			transformed_value = "X"
		case 0:
			transformed_value = " "
		}
		formatted_string[index] = transformed_value
	}
	return formatted_string
}

func (b *board) printBoard() {
	for i := 0; i < 3; i++ {
		to_print_values := transformValuesToStrings(b.board[i])
		fmt.Printf("| %v | %v | %v |\n", to_print_values[0], to_print_values[1], to_print_values[2])
	}
}

func (b *board) move() {
	fmt.Println("Please select the row and column of your movement in this format, row first then columns")
	var rowString, colString string
	fmt.Scanln(&rowString)
	fmt.Scanln(&colString)

	var rowInt, colInt int
	_, err := fmt.Sscan(rowString, &rowInt)
	if err != nil {
		panic(err)
	}
	_, err1 := fmt.Sscan(colString, &colInt)
	if err1 != nil {
		panic(err1)
	}
	b.board[rowInt-1][colInt-1] = b.turn
	if b.turn == 1 {
		b.turn = 2
	} else {
		b.turn = 1
	}

}

func (b *board) Build_game() {
	b.turn = 1
}
func (b *board) Start_game() {
	b.printBoard()
	for !b.over {
		b.move()
		b.printBoard()
	}
}
func main() {
	var game1 board
	game1.Build_game()
	game1.Start_game()
}

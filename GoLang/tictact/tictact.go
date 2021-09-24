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

func checkMul(array [3][3]int, configuration [2][2]int) (int, bool) {
	var n_in_pos [3]int
	current_pos := array[configuration[0][0]][configuration[0][1]]
	for i := 0; i < 3; i++ {
		n_in_pos[i] = current_pos
		configuration[0][0] += configuration[1][0]
		configuration[0][1] += configuration[1][1]
		current_pos = array[configuration[0][0]][configuration[0][1]]
	}
	if (n_in_pos[0] != 0) && (n_in_pos[0] == n_in_pos[1]) && (n_in_pos[0] == n_in_pos[2]) {
		return n_in_pos[0], true
	}
	return 0, false
}

func (b *board) check() (int, bool) {
	var win bool = false
	var winner int
	var config = [8][2][2]int{{{0, 0}, {0, 1}}, {{1, 0}, {0, 1}}, {{2, 0}, {0, 1}}, {{0, 0}, {1, 0}}, {{0, 1}, {1, 0}}, {{0, 2}, {1, 0}}, {{0, 0}, {1, 1}}, {{2, 0}, {-1, 1}}}
	for i := 0; i < 8 && win == false; i++ {
		winner, win = checkMul(b.board, config[i])
		if win {
			b.over = true
			break
		}
	}
	return winner, win
}

func (b *board) printBoard() {
	for i := 0; i < 3; i++ {
		to_print_values := transformValuesToStrings(b.board[i])
		fmt.Printf("| %v | %v | %v |\n", to_print_values[0], to_print_values[1], to_print_values[2])
	}
}

func (b *board) move(test bool) {
	if !test {
		fmt.Println("Please select the row and column of your movement in this format, row first then columns")
		var rowString, colString string
		fmt.Scanln(&rowString)
		fmt.Scanln(&colString)
	}

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

func (b *board) BuildGame() {
	b.turn = 1
}
func (b *board) StartGame() {
	b.printBoard()
	for !b.over {
		b.move()
		b.printBoard()
	}
}
func main() {
	var game1 board
	game1.BuildGame()
	game1.StartGame()
}

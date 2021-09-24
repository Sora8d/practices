package main

import (
	"fmt"
)

type board struct {
	player1     int
	player2     int
	board       [3][3]int
	over        bool
	winner      string
	turn        int
	number_play int
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
	position := configuration[0]
	increments := configuration[1]
	current_value := array[position[0]][position[1]]
	for i := 0; i < 3; i++ {
		n_in_pos[i] = current_value
		if position[0] != 2 {
			position[0] += increments[0]
		}
		if position[1] != 2 {
			position[1] += increments[1]
		}
		current_value = array[position[0]][position[1]]
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

func (b *board) move(external bool, moves [2]int) {
	var rowInt, colInt int
	if !external {
		fmt.Println("Please select the row and column of your movement in this format, row first then columns")
		var rowString, colString string
		fmt.Scanln(&rowString)
		fmt.Scanln(&colString)
		_, err := fmt.Sscan(rowString, &rowInt)
		if err != nil {
			panic(err)
		}
		_, err1 := fmt.Sscan(colString, &colInt)
		if err1 != nil {
			panic(err1)
		}
	} else {
		rowInt = moves[0]
		colInt = moves[1]
	}
	if (rowInt < 4 && colInt < 4) && b.board[rowInt-1][colInt-1] == 0 {
		b.board[rowInt-1][colInt-1] = b.turn
		if b.turn == 1 {
			b.turn = 2
		} else {
			b.turn = 1
		}
	} else {
		fmt.Println("That move is invalid")
	}

}

func (b *board) BuildGame() {
	b.turn = 1
}
func (b *board) StartGame(test bool, movearray [][2]int) {
	b.printBoard()
	b.number_play = 0
	if test {
		for !b.over {
			b.move(true, movearray[b.number_play])
			b.printBoard()
			b.check()
			b.number_play++
		}
	} else {
		for !b.over {
			b.move(false, [2]int{0, 0})
			b.printBoard()
			b.check()
			b.number_play++
		}
	}
	fmt.Println("Game Over")
}
func main() {
restart:
	var game1 board
	game1.BuildGame()
	test_moves := [6][2]int{
		{3, 1},
		{2, 1},
		{1, 2},
		{2, 2},
		{1, 3},
		{2, 3},
	}
	game1.StartGame(false, test_moves[:])
	if true {
		goto restart
	}
}

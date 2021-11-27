package game

import (
	"fmt"
	"ttt/src/utils"
)

var (
	board [9]string
	cursor int
	symb string
)

func render() {
	utils.Clrscr()
	fmt.Println("\n\n\n")

	var spaces string = GetSpaces(10)

	for i := 0; i < 9; i += 3 {
		fmt.Println(spaces + " ", board[i], "|", board[i + 1], "|", board[i + 2])
		if i < 6 {
			fmt.Println(spaces, "-----------")
		}
	}
}

func blinkCursor() {
	if board[cursor] == " " {
		board[cursor] = "_"
		return
	}
	board[cursor] = " "
}

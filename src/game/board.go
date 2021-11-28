package game

import (
	"github.com/nsf/termbox-go"
)

var (
	board  [9]rune
	cursor int
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

func render() {
	termbox.Clear(defaultColor, defaultColor)

	x, y := 10, 5

	renderBoard(x, y)
	renderBoardValues(x, y)

	termbox.Flush()
}

func renderBoard(_x, _y int) {
	x, y := _x, _y

	for i := 0; i < 3; i, y = i+1, y+1 {
		for j := 0; j < 2; j, x = j+1, x+4 {
			termbox.SetCell(x, y, '|', defaultColor, bgColor)
		}
		x = _x - 3
		if i == 2 {
			break
		}
		for j := 0; j < 11; j, x = j+1, x+1 {
			termbox.SetCell(x, y+1, '-', defaultColor, bgColor)
		}
		x, y = _x, y+1
	}
}

func renderBoardValues(_x, _y int) {
	x, y := _x-2, _y
	c := 0
	for i := 0; i < 3; i, y = i+1, y+1 {
		for j := 0; j < 3; j, x = j+1, x+4 {
			termbox.SetCell(x, y, board[c], defaultColor, bgColor)
			c++
		}
		x, y = _x-2, y+1
	}
}

func blinkCursor() {
	if board[cursor] == ' ' {
		board[cursor] = '_'
		return
	}
	board[cursor] = ' '
}

package game

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var (
	board  [9]rune
	cursor int
	moves  = 9
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

func render() {
	termbox.Clear(defaultColor, defaultColor)

	x, y := 10, 2

	renderBoard(x, y)
	renderBoardValues(x, y)
	renderTurnMsg()

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

func renderTurnMsg() {
	var msg string

	if isMyTurn {
		msg = "your turn"
	} else {
		msg = FriendName + "'s turn"
	}

	tbprint(8, 9, msg)

}

func gameOverPage(won bool) {
	termbox.Clear(defaultColor, defaultColor)
	var msg string

	if won {
		msg = "Congratulations you won!"
	} else {
		msg = "You lost " + FriendName + " won."
	}

	tbprint(2, 3, msg)
	termbox.Flush()
}

func gameTie() {
	termbox.Clear(defaultColor, defaultColor)
	tbprint(2, 3, "It's a tie. GG wellplayed!")
	termbox.Flush()
}

func tbprint(x, y int, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, defaultColor, bgColor)
		x += runewidth.RuneWidth(c)
	}
}

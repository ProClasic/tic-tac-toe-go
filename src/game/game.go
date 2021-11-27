package game

import (
	"github.com/nsf/termbox-go"
	"time"
	"ttt/src/utils"
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
)

func flickerCursor() {
	for {
		blinkCursor()
		time.Sleep(time.Millisecond * 300)
	}
}

func moveHorizontal(mag int) {
	for {
		cursor = cursor + mag
		validateCursor()
		if board[cursor] == " " {
			break
		}
	}
}

func moveVertical(mag int) {
	var prevCursor int = cursor
	cursor += 3 * mag
	validateCursor()
	if board[cursor] != " " {
		cursor = prevCursor
	}
}

func validateCursor() {
	if cursor > 8 {
		cursor -= 9
	} else if cursor < 0 {
		cursor += 9
	}
}

func StartGame(symb string) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer func() {
		termbox.Close()
		utils.Clrscr()
	}()

	for i := 0; i < 9; i++ {
		board[i] = " "
	}

	go listenToKeyboard(keyboardEventsChan)

	go flickerCursor()

	render()

mainloop:
	for {
		select {
		case e := <-keyboardEventsChan:
			switch e.eventType {
			case MOVE:
				board[cursor] = " "
				switch e.key {
				case termbox.KeyArrowRight:
					moveHorizontal(1)
				case termbox.KeyArrowDown:
					moveVertical(1)
				case termbox.KeyArrowLeft:
					moveHorizontal(-1)
				case termbox.KeyArrowUp:
					moveVertical(-1)
				}
			case INSERT:
				board[cursor] = symb
				moveHorizontal(1)
			case END:
				break mainloop
			}
		default:
			render()
		}
	}
}

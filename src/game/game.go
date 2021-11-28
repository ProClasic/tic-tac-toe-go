package game

import (
	"net"
	"strconv"
	"time"
	"ttt/src/utils"

	"github.com/nsf/termbox-go"
)

var (
	keyboardEventsChan = make(chan keyboardEvent)
	isMyTurn           = false
	symb               rune
)

func moveHorizontal(mag int) {
	for {
		cursor = cursor + mag
		validateCursor()
		if board[cursor] == ' ' {
			break
		}
	}
}

func moveVertical(mag int) {
	var prevCursor int = cursor
	cursor += 3 * mag
	validateCursor()
	if board[cursor] != ' ' {
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

func resumeMyTurn() {
	isMyTurn = true
	cursor = 0

	moveHorizontal(1)
	flickerCursor()
}

func flickerCursor() {
	for {
		if !isMyTurn {
			break
		}
		blinkCursor()
		time.Sleep(time.Millisecond * 300)
	}
}

func startGame(conn net.Conn) {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	termbox.HideCursor()
	defer func() {
		termbox.Close()
		utils.Clrscr()
	}()

	for i := 0; i < 9; i++ {
		board[i] = ' '
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
				if !isMyTurn {
					continue mainloop
				}
				board[cursor] = ' '
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
				if !isMyTurn {
					continue mainloop
				}
				board[cursor] = symb
				conn.Write([]byte(strconv.Itoa(cursor) + "e"))
				isMyTurn = false
				render()
				go receive(conn)
			case END:
				break mainloop
			}
		default:
			if isMyTurn {
				render()
				time.Sleep(time.Millisecond * 10)
			}
		}
	}
}

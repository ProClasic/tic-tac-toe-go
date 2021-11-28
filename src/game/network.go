package game

import (
	"bufio"
	"net"
	"strconv"
	"strings"

	"github.com/nsf/termbox-go"
)

func receive(conn net.Conn) {
	data, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		termbox.Close()
		panic(err)
	}

	data = strings.TrimSpace(data)
	if data == "lost" {
		gameOver = true
		gameOverPage(false)
		return
	}

	if pos, err := strconv.ParseInt(data, 10, 8); err != nil {
		panic(err)
	} else {
		if symb == 'O' {
			board[pos] = 'X'
		} else {
			board[pos] = 'O'
		}
	}
	moves--

	if checkTie() {
		gameTie()
		return
	}

	resumeMyTurn()
}

func getFriendName(conn net.Conn) string {
	data, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(data)
}

func sendUserName(conn net.Conn) {
	conn.Write([]byte(Username + "\n"))
}

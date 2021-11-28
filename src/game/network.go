package game

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

func receive(conn net.Conn) {
	data, err := bufio.NewReader(conn).ReadString('e')

	if err != nil {
		panic(err)
	}

	if pos, err := strconv.ParseInt(strings.TrimSpace(data[:1]), 10, 8); err != nil {
		panic(err)
	} else {
		if symb == 'O' {
			board[pos] = 'X'
		} else {
			board[pos] = 'O'
		}
	}

	resumeMyTurn()
}

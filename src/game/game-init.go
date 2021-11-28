package game

import (
	"ttt/src/tcp"
)

var (
	Username string
)

func Host() {
	conn := tcp.MakeHost()
	isMyTurn = true
	symb = 'X'
	startGame(conn)
}

func Join() {
	conn := tcp.MakeClient()
	symb = 'O'
	go receive(conn)
	startGame(conn)
}

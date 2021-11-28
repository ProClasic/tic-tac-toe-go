package game

import (
	"ttt/src/tcp"
)

var (
	Username   string
	FriendName string
)

func Host() {
	conn := tcp.MakeHost()
	isMyTurn = true
	symb = 'X'

	sendUserName(conn)
	FriendName = getFriendName(conn)
	startGame(conn)
}

func Join() {
	conn := tcp.MakeClient()
	symb = 'O'

	FriendName = getFriendName(conn)
	sendUserName(conn)
	go receive(conn)
	startGame(conn)
}

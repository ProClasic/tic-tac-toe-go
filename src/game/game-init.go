package game

import (
	"ttt/src/tcp"
)

var (
	Username string
)

func Host() {
	tcp.MakeServer()
}

func Join() {
	tcp.MakeClient()
}

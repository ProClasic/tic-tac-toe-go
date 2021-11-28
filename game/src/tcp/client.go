package tcp

import (
	"fmt"
	"net"
)

func MakeClient() net.Conn {
	fmt.Print("paste the game url: ")
	var url string
	fmt.Scanln(&url)
	conn, err := net.Dial("tcp", url)

	if err != nil {
		panic(err)
	}

	return conn
}

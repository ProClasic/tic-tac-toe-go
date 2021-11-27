package tcp

import (
	"fmt"
	"net"
)

func MakeHost() net.Conn {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println(`
			send this link to your friend localhost:8080
			waiting for other player connection...
	`)

	conn, err := ln.Accept()

	if err != nil {
		panic(err)
	}

	return conn
}

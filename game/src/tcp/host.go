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
waiting for connection. Make sure you are running ngrok server and your hotspot/wifi tethering is on.
	`)

	conn, err := ln.Accept()

	if err != nil {
		panic(err)
	}

	return conn
}

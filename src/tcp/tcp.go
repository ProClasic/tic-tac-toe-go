package tcp

import (
	"bufio"
	"net"
	"fmt"
	"os"
	// "strings"
	"sync"
)

func handleIncom(conn net.Conn, rChan chan string) {
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			panic(err)
		}

		rChan <- data
	}
}

func handleInp(iChan chan string) {
	for {
		data, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err)
		}

		iChan <- data
	}
}

func printMsg(rChan chan string) {
	for {
		r := <-rChan
		fmt.Print(r)
	}
}

func sendMsgS(conn net.Conn, iChan chan string) {
	for {
		data := <-iChan

		conn.Write([]byte(data))
	}
}

func MakeServer() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		rChan := make(chan string)
		iChan := make(chan string)

		go handleIncom(conn, rChan)
		go printMsg(rChan)
		go handleInp(iChan)
		go sendMsgS(conn, iChan)
	}
}

func sendMsg(conn net.Conn, iChan chan string, wg sync.WaitGroup) {
	for {
		data := <-iChan

		fmt.Fprintf(conn, data)
	}
	wg.Done()
}

func MakeClient() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	iChan := make(chan string)
	rChan := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)

	go handleIncom(conn, rChan)
	go printMsg(rChan)
	go handleInp(iChan)
	go sendMsg(conn, iChan, wg)

	wg.Wait()
}

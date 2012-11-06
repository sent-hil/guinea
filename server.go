package main

import (
	"fmt"
	"net"
)

type Hub struct {
	connections map[net.Conn]bool
	broadcast chan string
	register chan net.Conn
	unregister chan net.Conn
}

func main() {
	var hub = Hub {
		broadcast: make(chan string),
		register: make(chan net.Conn),
		unregister: make(chan net.Conn),
	}

	ln, _ := net.Listen("tcp", ":3000")

	// server loop
	for {
		conn, err := ln.Accept()
		if err != nil {
			// why continue?
			fmt.Println("ERR")
		}

		go func() { hub.register <- conn }()

		//conns := make(map[string]bool)
		//conns[conn.RemoteAddr().String()] = true
		//fmt.Println(conns)

		go handleClient(conn)

		select {
		case c := <- hub.register:
			fmt.Println(c)
		default:
			continue
		}

		// inside handleClient, once connect,
		// print remoteAddr

		// here read handleclient channel
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()

	// print client info
	//fmt.Println(conn.RemoteAddr())

	// move away / no use?
	// create a channel for message
	message := make(chan string)

	var buf [8]byte

	// client loop
	for {
		// read buffer
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// send buffer to channel
		go func() { message <- string(buf[0:n]) }()

		msg := <-message
		fmt.Println(msg)
	}
}

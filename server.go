package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			// why continue?
			fmt.Println("ERR")
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// close connection on exit
	defer conn.Close()
	
	// create a channel for message
	message := make(chan string)
	
	var buf [8]byte
	for {
		// read buffer
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		// send buffer to channel
		go func() {message <- string(buf[0:n])}()

		msg := <-message
		fmt.Println(msg)
	}
}
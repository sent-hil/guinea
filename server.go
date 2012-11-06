package main

import (
	"net"
	"fmt"
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
	
	var buf [8]byte
	for {
		// read buffer
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		
		// write buffer
		fmt.Println(buf)
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}
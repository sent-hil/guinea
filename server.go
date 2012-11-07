package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")

	go h.run()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR")
		}
		ncHandler(conn)
	}
}

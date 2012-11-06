package main

import (
	"net"
	"fmt"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")
	conn, err := ln.Accept()
	
	if err != nil {
		fmt.Println("ERR")
	}

	conn.Close()
}
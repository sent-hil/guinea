package main

import (
	"net"
	"fmt"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")

  message := make(chan string)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// why continue?
			fmt.Println("ERR")
		}
		go handleClient(conn, message)
	}
}

func handleClient(conn net.Conn, message chan string) {
	// close connection on exit
	defer conn.Close()
	
	var buf [8]byte
	for {
		// read buffer
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		//result, err2 := conn.Write(buf[0:n])

    // send buffer to channel
    message <- string(buf[0:n])

    select {
    case msg := <-message:
      fmt.Println(msg)
    default:
      fmt.Println("default")
    }
	}
}

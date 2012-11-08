
package main

import (
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":3000")

	// start hub
	go h.run()

	// loop to look for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR")
		}

		// handle individual connection
		go ncHandler(conn)
	}
}

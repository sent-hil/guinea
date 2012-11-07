package main

import "net"
import "fmt"

type connection struct {
	// no need to use a pointer since net.Conn is already a pointer
	nc   net.Conn
	send chan string
}

func (c *connection) reader() {
	for {
		// read message
		buf := make([]byte, 1)

		_, err := c.nc.Read(buf)

		if err != nil {
			break
		}

		// ignore enter/newline
		switch buf[0] {
		case 10:
		case 13:
		default:
			h.broadcast <- string(buf)
		}
	}
}

func (c *connection) writer() {
	for message := range c.send {
		// write message
		fmt.Println(message)
	}
	c.nc.Close()
}

func ncHandler(nc net.Conn) {
	// inits and sets connection
	c := &connection{send: make(chan string, 256), nc: nc}

	// registers connection
	h.register <- c

	// unregisters connection once handler exists
	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

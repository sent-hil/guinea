package main

import "net"
import "fmt"

type connection struct {
	nc   *net.Conn
	send chan string
}

func (c *connection) reader() {
	for {
		// read message
		buf := make([]byte, 8)

		_, err := (*c.nc).Read(buf)

		if err != nil {
			break
		}

		h.broadcast <- string(buf)
	}
}

func (c *connection) writer() {
	for message := range c.send {
		fmt.Println(message)
		// write message
	}
	(*c.nc).Close()
}

func ncHandler(nc *net.Conn) {
	// inits and sets connection
	c := &connection{send: make(chan string, 256), nc: nc}

	// registers connection
	h.register <- c

	// unregisters connection once handler exists
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

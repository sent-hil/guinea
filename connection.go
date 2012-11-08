package main

import (
	"net"
	//"fmt"
)

type connection struct {
	nc   net.Conn
	send chan []byte
}

func (c *connection) reader() {
	for {
		// read message
		buf := make([]byte, 8)

		_, err := c.nc.Read(buf)

		if err != nil {
			break
		}

		h.broadcast <- buf
	}
}

func (c *connection) writer() {
	for message := range c.send {
		// write message
		c.nc.Write(message)
	}
	c.nc.Close()
}

func ncHandler(nc net.Conn) {
	// inits and sets connection
	c := &connection{send: make(chan []byte, 256), nc: nc}

	// registers connection
	h.register <- c

	// unregisters connection once handler exists
	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

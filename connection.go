package main

import (
	"net"
	"io"
	//"fmt"
)

type connection struct {
	// type of connection
	ty io.ReadWriteCloser
	send chan []byte
}

type packet struct {
	conn *connection
	message []byte
}

func (c *connection) reader() {
	for {
		// read message
		buf := make([]byte, 8)

		_, err := c.ty.Read(buf)

		if err != nil {
			break
		}

		pkt := packet {
			conn: c,
			message: buf,
		}

		h.broadcast <- pkt
	}
}

func (c *connection) writer() {
	for message := range c.send {
		// write message
		c.ty.Write(message)
	}
	c.ty.Close()
}

func ncHandler(nc net.Conn) {
	// inits and sets connection
	c := &connection{send: make(chan []byte, 256), ty: nc}

	// registers connection
	h.register <- c

	// unregisters connection once handler exists
	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

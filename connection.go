package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type connection struct {
	// type of connection
	uid  string
	ty   io.ReadWriteCloser
	send chan []byte
}

type packet struct {
	conn    *connection
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

		pkt := packet{
			conn:    c,
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
	time_tos := time.Now().Unix()
	remote_tos := nc.RemoteAddr()
	uid := fmt.Sprintf("%d-%s", time_tos, remote_tos)

	// inits and sets connection
	c := &connection{send: make(chan []byte, 256), ty: nc, uid: uid}

	handler(c)
}

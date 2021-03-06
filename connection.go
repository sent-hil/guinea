package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"encoding/json"
)

type connection struct {
	// type of connection
	uid  string
	ty   io.ReadWriteCloser
	send chan json.RawMessage
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
	uid := get_uid(nc.RemoteAddr().String())

	// inits and sets connection
	c := &connection{send: make(chan json.RawMessage), ty: nc, uid: uid}

	handler(c)
}

// unique identifier for connection
func get_uid(remote_addr string) (uid string) {
	time_tos := time.Now().Unix()
	return fmt.Sprintf("%d-%s", time_tos, remote_addr)
}

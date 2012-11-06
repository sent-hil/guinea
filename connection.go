package main

import "net"
import "fmt"

type connection struct {
	nc *net.Conn
	send chan string
}

func (c *connection) reader() {
	for {
		var message string
		// receive message
		message = "foo"
		h.broadcast <- message
	}
	fmt.Println(c.nc)
}

func (c *connection) writer() {
	for message := range c.send {
		fmt.Println(message)
		// write message
	}
	//c.nc.Close()
}

func ncHandler(nc *net.Conn) {
	// inits and sets connection
	c := &connection{send: make(chan string, 256), nc :nc}

	// registers connection
	h.register <- c

	// unregisters connection once handler exists
	defer func() { h.unregister <- c}()
	go c.writer()
	c.reader()
}
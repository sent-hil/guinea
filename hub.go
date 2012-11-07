package main

import "fmt"

type hub struct {
	connections map[*connection]bool
	broadcast   chan string
	register    chan *connection
	unregister  chan *connection
}

var h = hub{
	connections: make(map[*connection]bool),
	broadcast:  make(chan string),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *hub) run() {
	// loops to look for incoming messages in channel
	for {
		select {
		case c := <-h.register:
			fmt.Println("registered: ", c)
		case c := <-h.unregister:
			fmt.Println("unregistered: ", c)
		case c := <-h.broadcast:
			fmt.Println("broadcast: ", c)
		}
	}
}

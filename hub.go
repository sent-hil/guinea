package main

import "fmt"

type hub struct {
	connections map[*connection]bool
	broadcast   chan []byte
	register    chan *connection
	unregister  chan *connection
}

var h = hub{
	connections: make(map[*connection]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *hub) run() {
	// loops to look for incoming messages in channel
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
			fmt.Println("registered: ", c)
		case c := <-h.unregister:
			fmt.Println("unregistered: ", c)
		case buf := <-h.broadcast:
			for key, _ := range h.connections {
				key.nc.Write(buf)
			}

			fmt.Println("broadcast: ", string(buf))
		}
	}
}

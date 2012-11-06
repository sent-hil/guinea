package main

import "fmt"

type hub struct {
	connections map[*connection]bool
	broadcast   chan string
	register    chan *connection
	unregister  chan *connection
}

var h = hub {
	broadcast:  make(chan string),
	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *hub) run() {
	for {
		select {
		case c:= <- h.register:
			fmt.Println("registered: ", c)
		}
	}
}

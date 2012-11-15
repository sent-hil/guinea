package main

import (
	"fmt"
	"encoding/json"
)

type hub struct {
	connections map[*connection]bool
	broadcast   chan packet
	register    chan *connection
	unregister  chan *connection
}

// TODO: change to more descriptive name, move to diff. place
var h = hub{
	connections: make(map[*connection]bool),
	broadcast:   make(chan packet),
	register:    make(chan *connection),
	unregister:  make(chan *connection),
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
		case pkt := <-h.broadcast:
			for conn, _ := range h.connections {
				if !(pkt.conn.uid == conn.uid) {
					mapD := map[string]string{conn.uid:string(pkt.message)}
					mapB, _ := json.Marshal(mapD)
					conn.send <- mapB
				}
			}

			fmt.Println("broadcast: ", string(pkt.message))
		}
	}
}

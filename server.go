package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net"
	"net/http"
	"text/template"
)

// our raw html
var homeTempl = template.Must(
	template.ParseFiles("resources/index.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
	// applies our template "to" req.Host (data obj)
	homeTempl.Execute(c, req.Host)
}

func main() {
	ln, _ := net.Listen("tcp", ":3000")
	go h.run()

	go func() {
		http.HandleFunc("/", homeHandler)
		http.Handle("/ws", websocket.Handler(wshandler))
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// loop to look for connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR")
		}

		// handle individual connection
		go ncHandler(conn)
	}
}

func wshandler(ws *websocket.Conn) {
	c := &connection{send: make(chan []byte, 256), ty: ws}

	handler(c)
}

func handler(c *connection) {
	h.register <- c

	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

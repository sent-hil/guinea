package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
)

// TODO: abstract the http handlers to higher order functions
var homeTempl = template.Must(template.ParseFiles("resources/index.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func handleHTTP() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/ws", websocket.Handler(wshandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleTCP() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("ERR")
		}

		// handle individual connection
		go ncHandler(conn)
	}
}

// TODO: let user specify which binary to run using flags
//			 rather having different binaries.
func main() {
	// TODO: call newHub(), which spawns goroutines inside
	go h.run()      // start hub
	go handleHTTP() // serve http requests
	handleTCP()     // listens and handles TCP connections

	// add quit channel later
}

func wshandler(ws *websocket.Conn) {
	uid := get_uid(ws.RemoteAddr().String())

	// TODO: abstract to func NewConnection() which takes
	//			 interface io.ReadWriteCloser()
	c := &connection{send: make(chan json.RawMessage), ty: ws,
		uid: uid}

	handler(c)
}

// TODO: change to more descriptive name
func handler(c *connection) {
	// TODO: change to method?
	h.register <- c

	//addr := strings.Split(c.uid, "-")
	//msg := []byte(fmt.Sprintf("%s %s", addr[1], "has joined."))

	//strB, _ := json.Marshal(msg)

	//h.broadcast <- packet{conn: c, message: strB}

	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

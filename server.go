package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

func page(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, index)
}

var index = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8" />
<script>
  var websocket = new WebSocket("ws://localhost:8080/ws");
  websocket.onmessage = function (msg) { console.log(msg); }
  websocket.onclose = function (msg) { console.log("FOO"); }
</script>
<body>
The Faith of Humanity Lives in The Black Liquid.
</body>
</html>
`

func main() {
	ln, _ := net.Listen("tcp", ":3000")
	go h.run()

	go func() {
		http.HandleFunc("/", page)
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

// TODO: DRY it out, it's very similar to ncHandler().
func wshandler(ws *websocket.Conn) {
	c := &connection{send: make(chan []byte, 256), ty: ws}

	h.register <- c

	defer func() { h.unregister <- c }()

	go c.writer()
	c.reader()
}

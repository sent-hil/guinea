package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"net"
	"os"
)

func handleInput(nc net.Conn) {
	// read from server async
	go func() {
			for {
				buf := make([]byte, 8)
				_, err := nc.Read(buf)

				if err != nil {
					break
				}

				fmt.Print(string(buf))
			}
	}()

	// write to server
	for {
		e := termbox.PollEvent()
		if e.Ch == 0 {
			switch e.Key {
			case termbox.KeyCtrlC:
				return
			}
		} else {
			// send to server
			nc.Write([]byte(string(e.Ch)))
			fmt.Print(string(e.Ch))
		}
	}
}

func main() {
	args := os.Args
	addr := args[1]

	// connects to tcp server
	nc, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	err2 := termbox.Init()
	if err2 != nil {
		panic(err2)
	}
	defer termbox.Close()

	handleInput(nc)
}

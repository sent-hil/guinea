package main

import (
	"fmt"
	termbox "github.com/nsf/termbox-go"
)

func handleInput() {
	for {
		e := termbox.PollEvent()
		if e.Ch == 0 {
			switch e.Key {
			case termbox.KeyCtrlC:
				return
			}
		} else {
			fmt.Print(string(e.Ch))
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	handleInput()
}
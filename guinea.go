package main

import (
  "net"
  "fmt"
  termbox "github.com/nsf/termbox-go"
)

func drawLine() {
  w, h := termbox.Size()
  for i := 0; i < w; i++ {
    termbox.SetCell(i, h-2, ' ', termbox.ColorBlack,
    termbox.ColorWhite)
  }
  termbox.Flush()
}

func handleInput(quit chan bool) {
  for {
    e := termbox.PollEvent()
    if e.Ch == 0 {
      switch e.Key {
      case termbox.KeyCtrlC:
        quit <- true
        return
      }
    }

    if e.Ch != 0 {
      fmt.Println("letter")
      return
    }
  }
}

func main() {
  ln, _ := net.Listen("tcp", ":3000")

  conn, _ := ln.Accept()

  err := termbox.Init()
  if err != nil {
    fmt.Println("ERROR:", err)
  }

  defer termbox.Close()

  //hist = []rune{}
  //go manager.Listen()

  quit := make(chan bool, 1)
  go handleInput(quit)

  <-quit

  conn.Close()
}

package main

import (
	"fmt"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	const coldef = termbox.ColorDefault
	termbox.Clear(coldef, coldef)
	termbox.Sync()
	termbox.SetInputMode(termbox.InputEsc)
	tbprint(1, 1, termbox.ColorWhite, termbox.ColorBlack, "Running...")
	termbox.Flush()

	time.Sleep(1 * time.Second)

	w, h := termbox.Size()
	msg := fmt.Sprintf("%5v %5v", w, h)
	tbprint(1, 1, termbox.ColorWhite, termbox.ColorBlack, msg)
	tbprint(1, 2, termbox.ColorWhite, termbox.ColorBlack, msg)
	termbox.Flush()
	// termbox.Sync()

	time.Sleep(3 * time.Second)

	msg = fmt.Sprintf("%5v %5v", 12, 0)
	tbprint(1, 1, termbox.ColorWhite, termbox.ColorBlack, msg)
	tbprint(1, 2, termbox.ColorWhite, termbox.ColorBlack, msg)
	termbox.Flush()
	// termbox.Sync()

	time.Sleep(3 * time.Second)

	// select {}
}

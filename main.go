package main

import (
	"log"
	"math/rand"

	"github.com/nsf/termbox-go"
)

const coldef = termbox.ColorDefault

func drawBox(x, y int) {
	termbox.Clear(coldef, coldef)
	termbox.SetCell(x, y, 'h', coldef, coldef)
	termbox.SetCell(x+1, y, 'h', coldef, coldef)
	termbox.SetCell(x, y+1, 'h', coldef, coldef)
	termbox.SetCell(x+1, y+1, 'h', coldef, coldef)
	termbox.Flush()
}

func main() {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	drawBox(0, 0)

MAINLOOP:
	for {
		w, h := termbox.Size()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break MAINLOOP
			}
		}
		drawBox(rand.Intn(w), rand.Intn(h))
	}
}

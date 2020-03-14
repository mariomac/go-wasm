package main

import (
	"time"

	"github.com/mariomac/go-wasm/internal/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	c.ClearRect(0,0,500,500)
	for {
		c.StrokeStyle("green")
		c.LineWidth(15)
		c.LineCap(draw.LineCapRound)
		c.StrokeRect(100,100,100,100)
		time.Sleep(1 * time.Millisecond)
	}

}

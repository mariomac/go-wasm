package main

import (
	"time"

	"github.com/mariomac/go-wasm/internal/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	c.ClearRect(0, 0, 500, 500)
	c.StrokeStyle("green")
	c.LineWidth(15)
	c.LineCap(draw.LineCapRound)
	for {
		c.Translate(500, 500)
		c.ClearRect(-500, -500, 1000, 1000)
		c.StrokeRect(-50, -50, 100, 100)
		c.Rotate(0.02)
		c.Translate(-500, -500)
		//c.Scale(1.001, 1.001)
		time.Sleep(20 * time.Millisecond)
	}

}

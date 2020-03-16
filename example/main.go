package main

import (
	"strings"
	"time"

	"github.com/mariomac/gorrazo/pkg/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	c.ClearRect(0, 0, c.GetWidth(), c.GetHeight())
	txtPos := float64(c.GetHeight())
	counter := float64(0)
	img := draw.HTMLImageElement("./img/virus.png")
	for {
		gradient := c.CreateLinearGradient(0, 0, c.GetWidth(), c.GetHeight())
		gradient.AddColorStop(1, "#ff00ff")
		gradient.AddColorStop(0.3, "#888888")
		gradient.AddColorStop(0, "#00ff88")
		c.FillStyleGradient(gradient)
		c.FillRect(0, 0, c.GetWidth(), c.GetHeight())

		c.FillStyle("yellow")
		c.StrokeStyle("black")
		c.Font("40px Arial")
		for l, line := range txt {
			c.FillText(line, 0, txtPos+float64(l)*50)
			c.StrokeText(line, 0, txtPos+float64(l)*50)
		}
		c.Save()
		c.ShadowBlur(15)
		c.ShadowColor("black")
		c.ShadowOffsetX(10)
		c.ShadowOffsetY(20)
		c.DrawImageD(img, 20, 20, 200, 200)
		c.DrawImage(img, 100, 100)
		c.Save()
		c.Translate(360, 160)
		c.Rotate(counter / 10)
		c.Translate(-360, -160)
		c.DrawImageD(img, 360, 160, 120, 120)
		c.Restore()
		c.Restore() // restore non-blur
		txtPos -= 1
		counter++
		time.Sleep(20 * time.Millisecond)
	}
}

var txt = strings.Split(`Canvas del Gorrazo
is a Go library
that allows you
operating your HTML
Canvas via Go type
safety thanks to
WebAssembly targetting.

Its name is inspired in
Cangas del Morrazo, a
Galician (Spain) city.
`, "\n")

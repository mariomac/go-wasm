package main

import (
	"time"

	"github.com/mariomac/go-wasm/internal/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	for {
		c.Clear("green")
		time.Sleep(1 * time.Millisecond)
	}

}

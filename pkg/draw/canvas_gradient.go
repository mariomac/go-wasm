package draw

import "syscall/js"

type CanvasGradient struct {
	value js.Value
}

func (c *CanvasGradient) AddColorStop(offset float64, color string) {
	c.value.Call("addColorStop", offset, color)
}

package draw

import "syscall/js"

type CanvasPattern struct {
	value js.Value
}

func (c *CanvasPattern) SetTransform(m DOMMatrixReadOnly) {
	c.value.Call("setTransform", m.value)
}
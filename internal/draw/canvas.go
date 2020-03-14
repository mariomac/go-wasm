package draw

import "syscall/js"

type Canvas struct {
	width  float64
	height float64
	ctx    js.Value
}

func GetCanvas(id string) *Canvas {
	doc := js.Global().Get("document")
	canvasElement := doc.Call("getElementById", id)

	canvas := Canvas{
		// todo: allow choosing full screen or custom definition
		width:  doc.Get("body").Get("clientWidth").Float(),
		height: doc.Get("body").Get("clientHeight").Float(),
		ctx:    canvasElement.Call("getContext", "2d"),
	}
	canvasElement.Set("width", canvas.width)
	canvasElement.Set("height", canvas.height)
	return &canvas
}

func (c *Canvas) Clear(bgcolor string) {
	c.ctx.Set("fillStyle", bgcolor)
	c.ctx.Call("fillRect", 0, 0, c.width, c.height)
}

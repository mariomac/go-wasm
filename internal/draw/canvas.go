package draw

import "syscall/js"

type Canvas struct {
	width  float64
	height float64
	doc    js.Value
	elem   js.Value
	ctx    js.Value
}

type CanvasConfig func(c *Canvas)

func GetCanvas(id string, cfgs ...CanvasConfig) *Canvas {
	doc := js.Global().Get("document")
	elem := doc.Call("getElementById", id)
	canvas := Canvas{
		doc:  doc,
		elem: elem,
		ctx:  elem.Call("getContext", "2d"),
	}

	for _, cfg := range cfgs {
		cfg(&canvas)
	}
	return &canvas
}

func FullScreen(resize bool) CanvasConfig {
	return func(c *Canvas) {
		c.adjustToWindow()
		if resize {
			js.Global().Set("onresize", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
				c.adjustToWindow()
				return nil
			}))
		}
	}
}

func (c *Canvas) adjustToWindow() {
	c.width = c.doc.Get("body").Get("clientWidth").Float()
	c.height = c.doc.Get("body").Get("clientHeight").Float()
	c.elem.Set("width", c.width)
	c.elem.Set("height", c.height)
}

func (c *Canvas) Clear(bgcolor string) {
	c.ctx.Set("fillStyle", bgcolor)
	c.ctx.Call("fillRect", 0, 0, c.width, c.height)
}

package draw

import (
	"syscall/js"
)

type Canvas struct {
	width  float64
	height float64
	doc    js.Value
	elem   js.Value
	ctx    js.Value
}

type Cap string

const (
	LineCapButt   Cap = "butt"
	LineCapRound  Cap = "round"
	LineCapSquare Cap = "square"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D
// TODO: measureText https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/measureText
// TODO: lineJoin https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineJoin
// TODO: miterLimit https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/miterLimit
// TODO: getLineDash, setLineDash, lineDashOffset https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/getLineDash
// TODO: textAlign, textBaseline, direction
// TODO: createLinearGradient, createRadialGradient, createPattern
// TODO: shadowBlur, shadowColor, shadowOffsetX, shadowOffsetY
// TODO: beginPath, closePath, moveTo, lineTo, bezierCurveTo, quadraticCurveTo, arc, arcTo, ellipse, rect
// TODO: fill, stroke, drawFocusIfNeeded, scrollPathIntoView, clip, isPointInPath, isPointInStroke
// TODO: currentTransform, getTransform
// TODO: globalAlpha, globalCompositeOperation
// TODO: drawImage
// TODO: createImageData, getImageData, putImageData
// TODO: imageSmoothingEnabled, imageSmoothingQuality
// TODO: save, restore, canvas
// TODO: addHitRegion, removeHitRegion, clearHitRegions
// TODO: filter

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

func (c *Canvas) GetWidth() float64 {
	return c.width
}
func (c *Canvas) GetHeight() float64 {
	return c.height
}

func (c *Canvas) FillRect(x, y, width, height float64) {
	c.ctx.Call("fillRect", x, y, width, height)
}

func (c *Canvas) ClearRect(x, y, width, height float64) {
	c.ctx.Call("clearRect", x, y, c.width, c.height)
}

func (c *Canvas) StrokeRect(x, y, width, height float64) {
	c.ctx.Call("strokeRect", x, y, width, height)
}

func (c *Canvas) FillText(text string, x, y float64) {
	c.ctx.Call("fillText", text, x, y)
}

func (c *Canvas) FillTextMaxWidth(text string, x, y, maxWidth float64) {
	c.ctx.Call("fillText", text, x, y, maxWidth)
}

func (c *Canvas) StrokeText(text string, x, y float64) {
	c.ctx.Call("strokeText", text, x, y)
}

func (c *Canvas) StrokeTextMaxWidth(text string, x, y, maxWidth float64) {
	c.ctx.Call("strokeText", text, x, y, maxWidth)
}

func (c *Canvas) LineWidth(width float64) {
	c.ctx.Set("lineWidth", width)
}

func (c *Canvas) LineCap(cap Cap) {
	c.ctx.Set("lineCap", string(cap))
}

func (c *Canvas) Font(font string) {
	c.ctx.Set("font", font)
}

func (c *Canvas) FillStyle(style string) {
	c.ctx.Set("fillStyle", style)
}

func (c *Canvas) StrokeStyle(style string) {
	c.ctx.Set("strokeStyle", style)
}

func (c *Canvas) Rotate(angleRadians float64) {
	c.ctx.Call("rotate", angleRadians)
}

func (c *Canvas) Scale(x, y float64) {
	c.ctx.Call("scale", x, y)
}

func (c *Canvas) Translate(x, y float64) {
	c.ctx.Call("translate", x, y)
}

func (ca *Canvas) Transform(a, b, c, d, e, f float64) {
	ca.ctx.Call("transform", a, b, c, d, e, f)
}

func (ca *Canvas) SetTransform(a, b, c, d, e, f float64) {
	ca.ctx.Call("setTransform", a, b, c, d, e, f)
}

func (ca *Canvas) SetTransformM(m DOMMatrixReadOnly) {
	ca.ctx.Call("setTransform", js.Value(m))
}

func (ca *Canvas) ResetTransform() {
	ca.ctx.Call("resetTransform")
}

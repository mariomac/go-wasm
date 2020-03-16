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

type Repetition string

const (
	Repeat   = "repeat"
	RepeatX  = "repeat-x"
	RepeatY  = "repeat-y"
	NoRepeat = "no-repeat"
)

// https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D
// TODO: measureText https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/measureText
// TODO: lineJoin https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/lineJoin
// TODO: miterLimit https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/miterLimit
// TODO: getLineDash, setLineDash, lineDashOffset https://developer.mozilla.org/en-US/docs/Web/API/CanvasRenderingContext2D/getLineDash
// TODO: textAlign, textBaseline, direction
// TODO: beginPath, closePath, moveTo, lineTo, bezierCurveTo, quadraticCurveTo, arc, arcTo, ellipse, rect
// TODO: fill, stroke, drawFocusIfNeeded, scrollPathIntoView, clip, isPointInPath, isPointInStroke
// TODO: currentTransform, getTransform
// TODO: globalAlpha, globalCompositeOperation
// TODO: drawImage
// TODO: createImageData, getImageData, putImageData
// TODO: imageSmoothingEnabled, imageSmoothingQuality
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

func (c *Canvas) FillStyleGradient(g CanvasGradient) {
	c.ctx.Set("fillStyle", g.value)
}

func (c *Canvas) FillStylePattern(p CanvasPattern) {
	c.ctx.Set("fillStyle", p.value)
}

func (c *Canvas) StrokeStyle(style string) {
	c.ctx.Set("strokeStyle", style)
}

func (c *Canvas) StrokeStyleGradient(g CanvasGradient) {
	c.ctx.Set("strokeStyle", g.value)
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
	ca.ctx.Call("setTransform", m.value)
}

func (ca *Canvas) ResetTransform() {
	ca.ctx.Call("resetTransform")
}

func (c *Canvas) Save() {
	c.ctx.Call("save")
}

func (c *Canvas) Restore() {
	c.ctx.Call("restore")
}

func (c *Canvas) Canvas() js.Value {
	return c.ctx
}

func (c *Canvas) DrawImage(img CanvasImageSource, dx, dy float64) {
	c.ctx.Call("drawImage", img.value, dx, dy)
}

func (c *Canvas) DrawImageD(img CanvasImageSource, dx, dy, dWidth, dHeight float64) {
	c.ctx.Call("drawImage", img.value, dx, dy, dWidth, dHeight)
}

func (c *Canvas) CreateLinearGradient(x0, y0, x1, y1 float64) CanvasGradient {
	grad := c.ctx.Call("createLinearGradient", x0, y0, x1, y1)
	return CanvasGradient{value: grad}
}

func (c *Canvas) CreateRadialGradient(x0, y0, r0, x1, y1, r1 float64) CanvasGradient {
	grad := c.ctx.Call("createRadialGradient", y0, y0, r0, x1, y1, r1)
	return CanvasGradient{value: grad}
}

func (c *Canvas) CreatePattern(img CanvasImageSource, r Repetition) CanvasPattern {
	pat := c.ctx.Call("createPattern", img.value, string(r))
	return CanvasPattern{value: pat}
}

func (c *Canvas) ShadowBlur(level float64) {
	c.ctx.Set("shadowBlur", level)
}

func (c *Canvas) ShadowColor(color string) {
	c.ctx.Set("shadowColor", color)
}

func (c *Canvas) ShadowOffsetX(offset float64) {
	c.ctx.Set("shadowOffsetX", offset)
}

func (c *Canvas) ShadowOffsetY(offset float64) {
	c.ctx.Set("shadowOffsetY", offset)
}

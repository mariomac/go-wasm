package draw

import "syscall/js"

type CanvasImageSource struct {
	value js.Value
}

// TODO: CSSImageValue, SVGImageElement, HTMLVideoElement, HTMLCanvasElement, ImageBitmap, OffscreenCanvas.

func HTMLImageElement(src string) CanvasImageSource {
	r := CanvasImageSource{
		value: js.Global().Get("Image").New(),
	}
	r.value.Set("src", src)
	return r
}

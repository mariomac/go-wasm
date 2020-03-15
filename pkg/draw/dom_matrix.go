package draw

import "syscall/js"

// At the moment, this seems too low level and not needed. Let's check later
type DOMMatrixReadOnly js.Value

type DOMMatrixFactory js.Value

func NewMatrixFactory() DOMMatrixFactory {
	return DOMMatrixFactory(js.Global().Get("DOMMatrixReadOnly"))
}

func (mf DOMMatrixFactory) ForValues(a, b, c, d, e, f float64) DOMMatrixReadOnly {
	return DOMMatrixReadOnly(js.Value(mf).New([]interface{}{a, b, c, d, e, f}))
}

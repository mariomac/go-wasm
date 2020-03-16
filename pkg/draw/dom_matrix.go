package draw

import "syscall/js"

// At the moment, this seems too low level and not needed. Let's check later
type DOMMatrixReadOnly struct {
	value js.Value
}

type DOMMatrixFactory js.Value

var constr = js.Global().Get("DOMMatrixReadOnly")

func NewDOMMatrixReadOnly(a, b, c, d, e, f float64) DOMMatrixReadOnly {
	return DOMMatrixReadOnly{
		value: constr.New([]interface{}{a, b, c, d, e, f}),
	}
}

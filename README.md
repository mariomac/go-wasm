Canvas del Gorrazo
==================

A Go library to manage HTML5 canvas through WebAssembly.

**This library is in early development stage**

## Example

```go
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/mariomac/gorrazo/pkg/draw"
)

func main() {
	c := draw.GetCanvas("theCanvas",
		draw.FullScreen(true))

	c.StrokeStyle("green")
	c.LineWidth(15)
	for {
    		c.ClearRect(0, 0, c.GetWidth(), c.GetHeight())
		c.StrokeRect(-50, -50, 100, 100)
		c.Rotate(0.02)
		time.Sleep(20 * time.Millisecond)
	}

```

How to run the bundled example:

```
cd examples
make clean build
```

Then load the `examples/site` folder from any local server.

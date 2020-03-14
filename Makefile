GOOS=js
GOARCH=wasm

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o site/main.wasm ./pkg

clean:
	@rm -f site/main.wasm


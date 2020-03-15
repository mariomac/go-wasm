GOOS=js
GOARCH=wasm

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o site/main.wasm

clean:
	@rm -f site/main.wasm


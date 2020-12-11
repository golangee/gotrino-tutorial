build:
	go generate ./...

lint:
	GOOS=js GOARCH=wasm go build ./...
	GOOS=js GOARCH=wasm golangci-lint run --enable-all --exclude-use-default=false --disable=paralleltest


run: build
	gotrino-make serve
build:
	go generate ./...

lint:
	GOOS=js GOARCH=wasm go build ./...
	GOOS=js GOARCH=wasm golangci-lint run --enable-all --exclude-use-default=false --disable=paralleltest

travis:
	GO111MODULE=off GOPROXY=direct go get -u github.com/golangee/gotrino-make/cmd/gotrino-make
	gotrino-make -dir=./dist build

run: build
	gotrino-make serve
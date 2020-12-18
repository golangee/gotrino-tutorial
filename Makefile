build:
	go generate ./...

lint:
	GOOS=js GOARCH=wasm go build ./...
	GOOS=js GOARCH=wasm golangci-lint run --enable-all --exclude-use-default=false --disable=paralleltest

travis:
	GO111MODULE=off GOPROXY=direct go get -u github.com/golangee/gotrino-make/cmd/gotrino-make
	gotrino-make -dir=./dist build
	gotrino-make -deploy-host=www527.your-server.de -deploy-user=$FTP_USER -deploy-password=$FTP_PWD  -deploy-src=./dist deploy-ftp

run: build
	gotrino-make serve
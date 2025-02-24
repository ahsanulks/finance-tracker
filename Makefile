.PHONY: test build

test:
	go test -v ./...

build:
	CGO_ENABLED=0 GOOS=$(shell go env GOOS) GOARCH=$(shell go env GOARCH) go build -o fintrack -ldflags="-s -w" . 

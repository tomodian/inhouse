.PHONY: install build test cli

install:
	go get github.com/mitchellh/gox
	go mod tidy

test:
	DEBUG=1 go test ./... -cover

build:
	cd cli && make build

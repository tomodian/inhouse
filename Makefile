.PHONY: install build test cli

install:
	go get github.com/mitchellh/gox
	go mod tidy

test:
	DEBUG=1 go test -cover -coverprofile=coverage.txt -covermode=atomic ./...

build:
	cd cli/inhouse && make build

.PHONY: install test cli

install:
	go mod tidy

test:
	DEBUG=1 go test ./... -cover

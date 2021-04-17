.PHONY: run test

CMD := "go run main.go flags.go run.go"

run:
	@echo
	@echo "Run:"
	@echo "  $(CMD)"
	@echo
	@echo "Debug:"
	@echo "  DEBUG=1 $(CMD)"
	@echo

test:
	DEBUG=1 go test ./...

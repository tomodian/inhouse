.PHONY: run test

APP   := inhouse
BUILD := ../build
CMD   := "go run main.go flags.go run.go"

run:
	@echo
	@echo "Run:"
	@echo "  $(CMD)"
	@echo
	@echo "Debug:"
	@echo "  DEBUG=1 $(CMD)"
	@echo

clean:
	mkdir -p $(BUILD)
	rm -Rf $(BUILD)/*

bundle-nix:
	cd $(BUILD) && find . -type f ! -name '*.exe' | xargs -I % sh -c "mv % $(APP) && zip %.zip $(APP) && rm -f $(APP)"

bundle-windows:
	cd $(BUILD) && find . -type f -name '*.exe' | xargs -I % sh -c "mv % $(APP).exe && zip %.zip $(APP).exe && rm -f $(APP).exe"

build: clean
	@echo "Building.."
	gox -output="$(BUILD)/$(APP)_{{.OS}}_{{.Arch}}" \
		-osarch="darwin/amd64" \
		-osarch="linux/arm" \
		-osarch="linux/amd64" \
		-osarch="windows/amd64"
	@echo "Bundling.."
	$(MAKE) bundle-nix
	$(MAKE) bundle-windows

test:
	DEBUG=1 go test ./...

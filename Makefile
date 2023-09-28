BIN=i3-status

.PHONY: default
default: build

.PHONY: help
help:
	@echo "Targets:"
	@echo "  build (default)"
	@echo "  clean"
	@echo "  install"
	@echo "  uninstall"
	@echo "  help"

.PHONY: build
build:
	go build -o $(BIN)

.PHONY: clean
clean:
	rm -f $(BIN)

.PHONY: install
install:
	mv $(BIN) /usr/local/bin

.PHONY: uninstall
uninstall:
	rm -f /usr/local/bin/$(BIN)

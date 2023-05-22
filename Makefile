BIN_DIR=$(shell pwd)/bin
BIN_EXISTS="$(shell ls ${BIN_DIR} | grep "libvoicevox_core.dylib" )"

.PHONY: build
build:
	chmod +x ./install_bin.bash
	./install_bin.bash
	go build

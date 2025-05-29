VERSION := $(shell git describe --tags --abbrev=0)
LDFLAGS := -X 'main.version=$(VERSION)'

.PHONY: build test clean

build:
	go build -ldflags="$(LDFLAGS)" -o bin/csv2json ./cmd/converter

test:
	go test -v -cover ./...

clean:
	rm -rf bin/
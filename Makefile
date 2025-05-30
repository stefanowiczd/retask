GOLANG_LINT_VERSION := $(shell golang-lint --version 2>/dev/null)

.PHONY: build gen run test mod-tidy mod-up lint all

all: require

require:
	@go version            >/dev/null 2>&1 || (echo "ERROR: Golang is required."; exit 1)
    @golang-lint --version >/dev/null 2>&1 || (echo "ERROR: golangci-lint is required."; exit 1)
	@mockgen --version     >/dev/null 2>&1 || (echo "ERROR: mockgen (https://github.com/uber-go/mock) is required."; exit 1)

build:
	go build .

run:
	go run .

lint:
	golangci-lint run

gen:
	go generate -v ./...

mod-tidy:
	go mod tidy

mod-up:
	go get -v -u ./...

test:
	go test -tags unit -v ./... -count=1


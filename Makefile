.PHONY: all build install lint race run test vendor

all: build test

# Build

build:
	go build ./...

install:
	go install ./...

lint:
	golint ./...

race:
	go test -race ./...

run:
	go run main.go

test:
	go test -cover ./...

vendor:
	go mod tidy
	go mod vendor

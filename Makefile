.PHONY: all build install lint race run test

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

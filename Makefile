BINARY_NAME=stanley

all: build test

build:
	go build -o bin/${BINARY_NAME} -race ./cmd/main.go

test:
	go test -race -v ./cmd/main.go
BINARY_NAME=stanley

all: build test

build-srv:
	go build -o ${BINARY_NAME} -race ./cmd/main.go

test:
	go test -race -v ./cmd/main.go
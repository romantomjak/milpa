.PHONY: build run test fmt

build:
	go build

run:
	./milpa

test:
	go test ./... -v

fmt:
	go fmt ./...

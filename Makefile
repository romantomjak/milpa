SHELL=zsh
PROJECT_ROOT := $(pathsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
GIT_COMMIT := $(shell git rev-parse HEAD)
GO_PKGS := $(shell go list ./...)

APP := milpa
VERSION := 1.0.0
PLATFORMS := darwin linux windows
os = $(word 1, $@)

.PHONY: build
build:
	go build -o ${APP} main.go

.PHONY: release
release: windows linux darwin

.PHONY: $(PLATFORMS)
$(PLATFORMS):
	@mkdir -p dist
	@GOOS=$(os) GOARCH=amd64 go build -o dist/$(os)/${APP} main.go
	@zip -q -X -j dist/${APP}_${VERSION}_$(os)_amd64.zip dist/$(os)/${APP}
	@rm -rf dist/$(os)

.PHONY: run
run:
	go run -race main.go

.PHONY: test
test:
	go test -cover ${GO_PKGS}

.PHONY: clean
clean:
	go clean

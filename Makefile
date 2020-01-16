.PHONY:frontend
GO           ?= go
GOFMT        ?= $(GO)fmt
FIRST_GOPATH := $(firstword $(subst :, ,$(shell $(GO) env GOPATH)))

pkgs          = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX       ?= $(shell pwd)
DIRNAME      ?= $(shell dirname $(shell pwd))

#TAG          ?= $(shell date +%s)
TAG          ?= $(shell git rev-parse --short HEAD)
TAGS		 ?= prod


RUN_ENV      ?= test

style:
	@echo ">> checking code style"
	@! $(GOFMT) -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

build:
	@echo ">> go build ..."
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags ${TAGS} --ldflags -w -o gormin main.go

swagger:
	@echo ">> swag init"
	@swag init

clean:
	@echo ">> remove gormin"
	@rm gormin

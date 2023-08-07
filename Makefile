GOPATH:=$(shell go env | grep GOPATH | sed 's/GOPATH=//' | tr -d '"')
GOBIN:=$(GOPATH)/bin

.PHONY: all
all: test lint

lint:
	$(GOBIN)/golangci-lint run

test:
	go clean -testcache
	go test ./... -race
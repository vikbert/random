SHELL := /bin/bash
export GOPROXY = https://proxy.golang.org

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$|(^#--)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m %-43s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m #-- /[33m/'

.PHONY: help
.DEFAULT_GOAL := help

BINARY = random
ifeq ($(OS),Windows_NT)
	BINARY := $(BINARY).exe
endif

#-- Golang commands
install: ## install the executable to GOPATH
	cd cmd/random && go install

tidy: ## verify the dependencies of module
	go mod tidy
	go mod verify

build: ## build the command "random"
	go build -race -o dist/$(BINARY) ./cmd/random

run: ## run the main.go
	go run ./cmd/random


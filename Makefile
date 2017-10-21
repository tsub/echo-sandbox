PROJECT  = echo-sandbox
SRC      ?= $(shell go list ./... | grep -v vendor)
TESTARGS ?= -v

deps:
	glide install
.PHONY: deps

build:
	go build -o build/$(PROJECT)
.PHONY: build

test:
	go test $(SRC) $(TESTARGS)
.PHONY: test

fmt:
	go fmt $(SRC)
.PHONY: fmt

vet:
	go vet $(SRC)
.PHONY: vet

dotenv:
	@echo "Create .env with .env.template"
	cp .env.template .env
.PHONY: dotenv

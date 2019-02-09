PROJECT     = echo-sandbox
GO111MODULE = on
SRC         ?= $(shell go list ./... | grep -v vendor)
TESTARGS    ?= -v

deps:
	go mod download
.PHONY: deps

build:
	go build -o build/$(PROJECT)
.PHONY: build

build_prod:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/$(PROJECT)
.PHONY: build_prod

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

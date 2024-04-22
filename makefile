SHELL := /bin/sh

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

## Docker up
up:
	docker-compose up
## Run tests
test: 
	$(GOCMD) test ./... -cover
## Generate wire_gen.go
wire: 
	cd internal/wire && wire
## Generate swagger docs
swag: 
	swag init -d ./cmd/api,./
## Unit test
mockgen:
	mockgen -source=internal/service/user_service.go -destination=internal/mock/user.go -package=mock
## Install dependencies
deps: 
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor
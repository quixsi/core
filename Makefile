.PHONY: build
build:
	@go build -o bin/server ./cmd/server
	@go build -o bin/cli ./cmd/cli

.PHONY: dev/server
dev/server:
	@go run ./cmd/server

.PHONY: dev/cli
dev/cli:
	@go run ./cmd/cli

.PHONY: test
test:
	@go test -v ./...

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: mod
mod:
	@go mod tidy
	@go mod verify

.PHONY: clean
clean:
	@rm -rf bin/

.PHONY: clean/all
clean/all: clean
	@go clean -modcache

.PHONY: deps
deps:
	@go mod download

# Include UI Makefile with ui/ prefix
ui/%:
	@make -C ui $(*)

# Import all UI targets
include ui/Makefile
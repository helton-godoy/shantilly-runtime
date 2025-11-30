# Shantilly Runtime - Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=shantilly-runtime
BINARY_UNIX=$(BINARY_NAME)_unix

# Build settings
LDFLAGS=-ldflags "-X main.version=$(shell git describe --tags --always) -X main.commit=$(shell git rev-parse --short HEAD) -X main.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')"

.PHONY: all build clean test coverage deps help

all: test build

build: 
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v

test:
	$(GOTEST) -v -race -coverprofile=coverage.out ./...

coverage: test
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f coverage.out coverage.html

deps:
	$(GOMOD) download
	$(GOMOD) tidy

run:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v 
	./$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BINARY_UNIX) -v

# Development tools
fmt:
	$(GOCMD) fmt ./...

vet:
	$(GOCMD) vet ./...

lint:
	golangci-lint run

# CI/CD helpers
ci-test:
	$(GOTEST) -v -race -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -func=coverage.out

ci-build:
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v

# Docker
docker-build:
	docker build -t shantilly-runtime:latest .

docker-run:
	docker run --rm -it shantilly-runtime:latest

# Help
help:
	@echo "Available targets:"
	@echo "  all        - Run tests and build"
	@echo "  build      - Build the binary"
	@echo "  test       - Run tests"
	@echo "  coverage   - Run tests with coverage report"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Download and tidy dependencies"
	@echo "  run        - Build and run the application"
	@echo "  fmt        - Format Go code"
	@echo "  vet        - Run go vet"
	@echo "  lint       - Run golangci-lint"
	@echo "  build-linux- Cross compile for Linux"
	@echo "  docker-build- Build Docker image"
	@echo "  docker-run  - Run Docker container"
	@echo "  help       - Show this help message"

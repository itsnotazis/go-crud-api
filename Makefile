# Variables
APP_NAME = go-crud-api
BUILD_DIR = build
CMD = cmd

# Default Go commands
GO = go
GOFMT = gofmt
GOTEST = go test
GOLINT = golangci-lint run

# Default target
.PHONY: all
all: run

# Run the application
.PHONY: run
run:
	$(GO) run $(CMD)/main.go

# Build the application
.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(CMD)/main.go

# Test the application
.PHONY: test
test:
	$(GOTEST) ./...

# Format the code
.PHONY: fmt
fmt:
	$(GOFMT) -w .

# Lint the code (install golangci-lint first)
.PHONY: lint
lint:
	$(GOLINT)

# Clean build files
.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

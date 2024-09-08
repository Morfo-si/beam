# Define variables for the application
APP_NAME := beam
SRC := main.go
ENV_FILE := .env

# Default Go commands
GO := go
GOBUILD := $(GO) build
GORUN := $(GO) run
GOTEST := $(GO) test
GOCLEAN := $(GO) clean
GOMOD := $(GO) mod

# Rule to build the Go binary
.PHONY: build
build:
	@echo "Building the application..."
	$(GOBUILD) -o $(APP_NAME) $(SRC)

# Rule to run the Go application
.PHONY: run
run:
	@echo "Running the application with environment variables from $(ENV_FILE)..."
	@if [ -f $(ENV_FILE) ]; then \
		export $$(cat $(ENV_FILE) | xargs); \
	fi; \
	$(GORUN) $(SRC)

# Rule to run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GOTEST) ./...

# Rule to format Go code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Rule to clean up Go binaries and cache
.PHONY: clean
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	rm -f $(APP_NAME)

# Rule to tidy up Go dependencies
.PHONY: tidy
tidy:
	@echo "Tidying Go modules..."
	$(GOMOD) tidy

# Rule to vendor Go dependencies
.PHONY: vendor
vendor:
	@echo "Vendoring Go modules..."
	$(GOMOD) vendor

# Rule to install dependencies
.PHONY: deps
deps:
	@echo "Installing Go dependencies..."
	$(GO) get -v ./...

# Rule to display help message
.PHONY: help
help:
	@echo "Makefile commands for Go project:"
	@echo "  build     - Build the Go application"
	@echo "  run       - Run the Go application"
	@echo "  test      - Run tests"
	@echo "  fmt       - Format the Go code"
	@echo "  clean     - Clean up built files"
	@echo "  tidy      - Tidy Go modules"
	@echo "  vendor    - Vendor Go dependencies"
	@echo "  deps      - Install Go dependencies"
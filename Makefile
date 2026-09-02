NAME=mirrorbot
MAIN_PATH=./cmd/mirrorbot

GO=go
LINTER=golangci-lint

.PHONY: all build run test lint fmt vet tidy clean help

all: build

build:
	@echo "> Compiling $(NAME)..."
	$(GO) build -o $(NAME) $(MAIN_PATH)

run: build
	@echo "> Running $(NAME) with test environment..."
	@GITHUB_EVENT_NAME="issues" \
	GITHUB_EVENT_PATH="./testdata/issue_opened.json" \
	INPUT_GITLAB_TOKEN="your_test_token" \
	INPUT_GITLAB_PROJECT_ID="12345678" \
	./$(NAME)

test:
	@echo "> Running tests..."
	$(GO) test -v -race ./...

fmt:
	@echo "> Formatting code..."
	$(GO) fmt ./...

lint: fmt
	@echo "> Linting code..."
	$(LINTER) run ./...

vet:
	@echo "> Verifying code..."
	$(GO) vet ./...

tidy:
	@echo "> Cleaning dependencies..."
	$(GO) mod tidy

clean:
	@echo "> Cleaning artifacts..."
	$(GO) clean
	rm -f $(NAME)

help:
	@echo "Available Makefile commands:"
	@grep -E '^[a-zA-Z_-]+:' $(MAKEFILE_LIST) | sed 's/://' | grep -v 'PHONY'
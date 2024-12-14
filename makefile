.PHONY: test fmt run help

test:
	@echo "Running tests..."
	@go test -v -count=1 -timeout 30s ./...

fmt:
	@echo "Formatting Go files"
	@find . -name '*.go' -exec gofmt -s -w {} \;

run:
	@echo "Running solvePartOne..."
	@go run day01/code.go

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  test    Run tests"
	@echo "  fmt     Format Go files"
	@echo "  run     Run solvePartOne"
	@echo "  help    Show this help message"
.DEFAULT_GOAL := help
help:
	@echo "Available targets:"
	@echo ""
	@echo "  test       Run the tests"
	@echo "  fmt        Format all Go files"
	@echo "  vet        Lint all Go files"
	@echo "  help       Show this list of available targets"


# Run the tests
test: fmt vet
	go test ./...


# Format all Go files
fmt:
	go fmt ./...

# Lint all Go files
vet:
	go vet ./...


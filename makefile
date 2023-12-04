.PHONY: build
build:
	@echo "Building the application..."
	go build -v ./cmd/app
	@echo "Build complete."

.PHONY: test
test: 
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

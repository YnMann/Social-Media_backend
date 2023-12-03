.PHONY: build
build:
	@echo "Building the application..."
	go build -v ./cmd/app
	@echo "Build complete."

.DEFAULT_GOAL := build

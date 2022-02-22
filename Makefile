.PHONY: all
all:
	@echo "Starting make"
	@go mod tidy
	@go fmt ./...
	@go install -v -ldflags="-s -w" ./cmd/rs

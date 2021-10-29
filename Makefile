.PHONY: all
all:
	@echo "Starting make"
	@go fmt ./...
	@go mod tidy
	@go install -v

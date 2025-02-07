include .env
include .env

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help Show this help"
	@echo "  setup Setup"
	@echo "  start Start Server"
	@echo "  start-dev Watch changes"
	@echo "  lint Lint code"

setup:
	@go install go.uber.org/nilaway/cmd/nilaway@latest
start:
	@echo "Start server..."
	@go run main.go	
start-dev:
	@echo "Start server and live reload..."
	@air
lint:
	@echo "Linting..."
	@golangci-lint run

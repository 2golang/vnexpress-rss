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
	@echo "  install_packages Install packages"
	@echo "  update_packages Update packages"
	@echo "  generate_docs Generate docs"

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
install_packages:
	@echo "Installing packages..."
	@go mod tidy
update_packages:
	@echo "Updating packages..."
	@go get -u
	@go mod tidy
generate_docs:
	@echo "Generating docs..."
	@swag init -g main.go
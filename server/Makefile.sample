include .env

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  help                    Show this help"
	@echo "  setup                   Setup development environment"
	@echo "  start                   Start Server"
	@echo "  build                   Build Server"
	@echo "  start-dev               Watch changes with live reload"
	@echo "  lint                    Lint code"
	@echo "  install_packages        Install packages"
	@echo "  update_packages         Update packages"
	@echo "  generate_docs          Generate Swagger docs"
	@echo "  db-create-migration    Create migration. Usage: make db-create-migration name=<migration_name>"
	@echo "  db-command             Run database commands with goose"
	@echo "  db-status              Show current migration status"
	@echo "  db-up                  Run all pending migrations"
	@echo "  db-down                Rollback last migration"
	@echo "  db-reset               Reset database to initial state"

setup:
	@go install go.uber.org/nilaway/cmd/nilaway@latest
start:
	@echo "Start server..."
	@go run main.go
build:
	@echo "Building..."
	@ls -l
	@pwd
	@go build -o . main.go
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
db-command:
	@echo "Running db command..."
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="postgres://postgres:123456@localhost:5432/golang_rss" \
	GOOSE_MIGRATION_DIR=./db/migrations \
	goose $(cmd)

db-create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: name parameter is required. Usage: make db-create-migration name=<migration_name>"; \
		exit 1; \
	fi
	@echo "Creating migration..."
	@$(MAKE) db-command cmd="create $(name) sql"

db-status:
	@echo "Checking migration status..."
	@$(MAKE) db-command cmd="status"
db-up:
	@echo "Migrating up..."
	@$(MAKE) db-command cmd="up"
db-down:
	@echo "Migrating down..."
	@$(MAKE) db-command cmd="down"
db-reset:
	@echo "Resetting database..."
	@$(MAKE) db-command cmd="reset"
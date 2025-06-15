.PHONY: up down build test test-coverage logs clean swagger sqlc migrate migrate-down migrate-create tools

# Команды для Docker
up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

logs:
	docker-compose logs -f logs-api

# Команды для разработки
dev:
	air

install:
	go mod download
	go mod tidy

# Команды для работы с БД
sqlc:
	sqlc generate

# Todo Remove
migrate:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

# Todo Remove
migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

# Создание миграции (для Windows PowerShell)
# Использование: make migrate-create name=my_migration
migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

# Установка инструментов
tools:
	go install github.com/air-verse/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# Команды для тестирования
test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Swagger документация
swagger:
	swag init

# Очистка
clean:
	rm -f coverage.out coverage.html
	go clean
	docker-compose down -v --remove-orphans 
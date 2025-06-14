.PHONY: up down build test test-coverage logs clean

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

# Команды для тестирования
test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Очистка
clean:
	rm -f coverage.out coverage.html
	go clean
	docker-compose down -v --remove-orphans 
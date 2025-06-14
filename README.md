# Logs API

Простое REST API приложение на Go для записи логов через Kafka topic.

## 📋 Требования

- Go 1.24 или выше
- Docker и Docker Compose
- Make (опционально)

## 🚀 Быстрый старт

### Локальная разработка с Docker

```bash
# Запуск всех сервисов
make up

# Остановка всех сервисов
make down
```

### Локальная разработка без Docker

```bash
# Установка зависимостей
go mod download

# Запуск приложения
go run main.go
```

## 🛠 Команды Make

- `make up` - запуск всех сервисов в Docker
- `make down` - остановка всех сервисов
- `make build` - сборка Docker образа
- `make logs` - просмотр логов приложения
- `make test` - запуск тестов
- `make test-coverage` - запуск тестов с покрытием

## 🧪 Тестирование

```bash
# Запуск всех тестов
make test

# Запуск тестов с покрытием
make test-coverage
```

## 📁 Структура проекта

```
.
├── .air.toml          # Конфигурация для hot-reload
├── Dockerfile         # Dockerfile для сборки приложения
├── docker-compose.yml # Docker Compose конфигурация
├── internal/          # Внутренние пакеты приложения
├── main.go           # Точка входа в приложение
└── Makefile          # Команды для упрощения разработки
```

## 📝 API Документация

[TODO: Добавить описание API endpoints]


# Logs API

Простое REST API приложение на Go для записи логов в PostgreSQL.

## 📋 Требования

- Go 1.24 или выше
- Docker и Docker Compose
- Make (опционально)
- PostgreSQL 17.5 или выше

## 🚀 Быстрый старт

### Локальная разработка с Docker

```bash
# Запуск всех сервисов
make up

# Генерация кода для работы с БД
make sqlc

# Применение миграций
make migrate

# Остановка всех сервисов
make down
```

### Локальная разработка без Docker

```bash
# Установка зависимостей
go mod download

# Установка инструментов
make tools

# Генерация кода для работы с БД
make sqlc

# Применение миграций
make migrate

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
- `make sqlc` - генерация кода для работы с БД
- `make migrate` - применение миграций
- `make tools` - установка необходимых инструментов

## 🗄️ База данных

### Миграции

Для управления схемой базы данных используется [golang-migrate/migrate](https://github.com/golang-migrate/migrate).

```bash
# Создание новой миграции
migrate create -ext sql -dir migrations -seq название_миграции

# Применение миграций
make migrate
```

### Генерация кода

Для генерации типобезопасного кода для работы с БД используется [sqlc](https://sqlc.dev/).

```bash
# Генерация кода
make sqlc
```

### Структура таблицы logs

| Поле       | Тип                     | Описание                    |
|------------|-------------------------|----------------------------|
| id         | SERIAL PRIMARY KEY      | Уникальный идентификатор   |
| model      | VARCHAR(255)            | Название модели            |
| user_uuid  | UUID                    | UUID пользователя          |
| user_name  | VARCHAR(255)            | Имя пользователя           |
| action     | VARCHAR(100)            | Действие                   |
| action_at  | TIMESTAMP WITH TIME ZONE| Время действия             |
| description| TEXT                    | Описание                   |
| created_at | TIMESTAMP WITH TIME ZONE| Время создания записи      |

## 📁 Структура проекта

```
.
├── .air.toml           # Конфигурация для hot-reload
├── Dockerfile          # Dockerfile для сборки приложения
├── docker-compose.yml  # Docker Compose конфигурация
├── migrations/         # Файлы миграций
├── sql/               # SQL запросы
│   └── queries/       # Запросы для sqlc
├── internal/          # Внутренние пакеты приложения
│   ├── api/          # API handlers
│   ├── config/       # Конфигурация
│   ├── database/     # Работа с БД
│   └── db/          # Сгенерированный код sqlc
├── main.go           # Точка входа в приложение
├── sqlc.yaml         # Конфигурация sqlc
└── Makefile          # Команды для упрощения разработки
```

## 📝 API Документация

### Swagger

Для генерации Swagger документации используется [swaggo/swag](https://github.com/swaggo/swag).

#### Локальная установка установка swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

#### Генерация документации

```bash
# Генерация Swagger спецификации
swag init

# Или с помощью Make
make swagger
```

После генерации документация будет доступна по адресу:
- Swagger UI: `http://localhost:8080/swagger/index.html`
- JSON спецификация: `http://localhost:8080/swagger/doc.json`

## TODO

- [x] ~~Инициализировать просто http приложение~~
- [x] ~~Генерация swagger документации~~
- [x] ~~Реализация миграции базы данных и сущость Log~~
- [ ] !!! Сделать чтобы swagger документация обновлялась без перезапуска docker контейнера
- [ ] Слушателя кафки
- [ ] Добавление логов в базу данных
- [ ] Ендпионт получения логов

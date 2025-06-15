FROM golang:1.24-alpine

# Установка необходимых инструментов
RUN apk add --no-cache make git

WORKDIR /app

# Копируем файлы для сборки
COPY go.mod go.sum ./
COPY Makefile ./

# Установка зависимостей и инструментов
RUN make install && \
    make tools

# Копируем исходный код
COPY . .

# Генерация кода и документации
#RUN make sqlc && \
#    make swagger

# Команда по умолчанию для разработки
CMD ["air", "-c", ".air.toml"] 
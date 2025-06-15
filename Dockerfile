FROM golang:1.24-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download && \
go install github.com/swaggo/swag/cmd/swag@latest

CMD ["air", "-c", ".air.toml"]
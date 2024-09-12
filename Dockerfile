# Этап сборки (Build Stage)
FROM golang:1.22-alpine AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY backend/go.mod backend/go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY backend/ .

# Сборка бинарного файла
RUN go build -o main ./cmd/main.go

# Этап выполнения (Runtime Stage)
FROM alpine:latest

# Открываем порт 8080
EXPOSE 8080

# Создаем директорию для приложения
RUN mkdir /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=build /app/main /app/main

# Указываем команду для запуска приложения
ENTRYPOINT ["/app/main"]

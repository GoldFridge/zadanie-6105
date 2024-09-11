# Используем образ для сборки приложения
FROM golang:1.20-alpine AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY . .

# Переходим в папку backend и собираем бинарный файл
WORKDIR /app/backend
RUN go build -o /app/main ./cmd/main.go

# Минимальный образ для запуска приложения
FROM alpine:latest

# Создаем директорию для приложения
RUN mkdir /app

# Копируем скомпилированное приложение из предыдущего шага
COPY --from=build /app/main /app/main

# Указываем порт, на котором будет работать приложение
EXPOSE 8080

# Указываем команду для запуска приложения
ENTRYPOINT ["/app/main"]

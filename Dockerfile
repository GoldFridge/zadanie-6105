# Используем образ для сборки приложения
FROM golang:1.20-alpine AS build

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum в рабочую директорию
COPY backend/go.mod backend/go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь исходный код в рабочую директорию
COPY backend/ .

# Сборка бинарного файла
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

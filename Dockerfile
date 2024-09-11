# Указываем базовый образ Go для сборки
FROM golang:1.20 as builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для того, чтобы можно было установить зависимости
COPY backend/go.mod backend/go.sum ./backend/

# Переходим в директорию backend
WORKDIR /app/backend

# Устанавливаем зависимости
RUN go mod download

# Копируем все исходные файлы в контейнер
COPY backend/ .

# Собираем бинарный файл
RUN go build -o /app/main ./cmd/main.go

# Используем минимальный образ для запуска
FROM gcr.io/distroless/base-debian10

# Указываем рабочую директорию
WORKDIR /app

# Копируем бинарный файл из предыдущего шага
COPY --from=builder /app/main .

# Указываем команду для запуска приложения
CMD ["/app/main"]

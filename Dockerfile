# data-preparer/Dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o datapreparer ./cmd/datapreparer

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем бинарник
COPY --from=builder /app/datapreparer .

# Копируем конфиги (если нужно)
COPY ./config ./config

CMD ["./datapreparer"]
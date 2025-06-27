FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/html ./html
COPY --from=builder /app/decks ./decks
COPY --from=builder /app/app .

EXPOSE 8080

ENV PORT=8080

CMD ["./app"]
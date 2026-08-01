FROM golang:1.26.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./booking-service/main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 50053

ENTRYPOINT ["./app"]
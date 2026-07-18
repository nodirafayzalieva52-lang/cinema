FROM golang:1.26.5_alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
PUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app ./main.go

FROM scratch

COPY --from=builder /app/app /app

EXPOSE 8080

ENTRYPOINT ["/app"]

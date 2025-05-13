FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/main .

EXPOSE 8000

CMD ["/app/main"]
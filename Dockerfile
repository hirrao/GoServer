FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8121

CMD ["./app"]
FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o bin /app/cmd/server/main.go

EXPOSE 50051

ENTRYPOINT ["/app/bin"]
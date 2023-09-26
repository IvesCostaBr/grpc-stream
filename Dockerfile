FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go build -o bin /app/cmd/server/main.go

FROM grc.io/distroless/base-debian10

WORKDIR /build

EXPOSE 50051

COPY --from=builder /app/bin /build/bin

ENTRYPOINT ["/build/bin"]
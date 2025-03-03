FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:linux 

WORKDIR /app

COPY --from=builder /app/main .

COPY  .env .

EXPOSE 8080

CMD ["./main"]
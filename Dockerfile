FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY ./config/.env /app/config/.env


RUN go build -o main ./cmd/server  

EXPOSE 8080

CMD ["./main"]
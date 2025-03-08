FROM golang:1.23-bookworm AS base

WORKDIR /usr/app

COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o bin/main cmd/main.go

EXPOSE 8000

CMD ["./bin/main"]
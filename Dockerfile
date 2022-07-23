# syntax=docker/dockerfile:1

FROM golang:1.18
WORKDIR /app
COPY . ./
RUN go mod download
RUN go mod tidy

RUN go build -o bin/start_server ./cmd/server/main.go

CMD ["./bin/start_server"]
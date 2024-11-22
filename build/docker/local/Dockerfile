FROM golang:1.22.2-alpine3.18

RUN apk update && apk upgrade && \
    apk add --no-cache bash

WORKDIR /app

# Install dependencies first so we cache this layer and accelerate build times

RUN go install github.com/cosmtrek/air@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compile the app
RUN go build -v -o main ./cmd/api

EXPOSE 8080

# Run the bin
CMD ["air", "-c", ".air.toml"]
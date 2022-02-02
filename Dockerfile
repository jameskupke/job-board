FROM golang:1.17 as build
WORKDIR /app
COPY . .
CMD go run .

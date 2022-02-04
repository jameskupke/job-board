FROM golang:1.17 as build
RUN go install github.com/cortesi/modd/cmd/modd@latest
WORKDIR /app
COPY . .
CMD go run .

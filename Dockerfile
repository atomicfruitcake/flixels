FROM golang:latest

LABEL maintainer="atomicfruitcake"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build ./srv/main.go   

EXPOSE 8080

CMD ["./main"]
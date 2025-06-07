FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN go mod download

RUN go build -o app

ENTRYPOINT ["./app"]

FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app .

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080

ENTRYPOINT ["./app"]

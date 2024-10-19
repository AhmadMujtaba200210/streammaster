#build configurations
FROM golang:1.23.2 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/streammaster cmd/main.go

#runtime configurations
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /out/streammaster .

COPY --from=builder /app/config /root/config

EXPOSE 8080

CMD ["./streammaster"]

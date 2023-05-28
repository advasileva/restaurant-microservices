FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o auth cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/auth /build/auth

CMD ["./auth"]
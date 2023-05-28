FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o order cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/order /build/order

CMD ["./order"]
FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o service cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/service /build/service
COPY --from=builder /build/config/config.json /build/config/config.json

CMD ["./service"]
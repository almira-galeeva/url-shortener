FROM golang:1.18.3-alpine AS builder

COPY . /github.com/almira-galeeva/url-shortener/
WORKDIR /github.com/almira-galeeva/url-shortener/

RUN go mod download
RUN go build -o ./bin/url_shortener cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/almira-galeeva/url-shortener/bin/url_shortener .
COPY --from=builder /github.com/almira-galeeva/url-shortener/config/ /root/config/

EXPOSE 50051
EXPOSE 8080

CMD ["./url_shortener"]
FROM golang:1.15.2 AS builder

WORKDIR /quote
COPY . .

RUN CGO_ENABLED=0 go install -v ./...

FROM alpine:3.8


COPY --from=builder /go/bin/quote /quote
COPY --from=builder /quote/config.yaml /config.yaml

EXPOSE 9000
ENTRYPOINT ["/quote"]

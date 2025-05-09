FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

RUN apk add --no-cache git=2.47.2-r0 make=4.4.1-r2

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build


FROM alpine:3.21 AS deploy

RUN apk add --no-cache ca-certificates=20241121-r1 && update-ca-certificates

RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app

COPY --from=builder /app/hivebox .
COPY config/config.yaml ./config/config.yaml

EXPOSE 8888

ENTRYPOINT ["./hivebox"]


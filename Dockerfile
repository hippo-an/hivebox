FROM golang:1.23.4-bookworm as builder

WORKDIR /app

COPY . .

RUN go build -o hivebox ./*.go



FROM ubuntu:24.04 as deploy

WORKDIR /app

COPY --from=builder /app/hivebox .

EXPOSE 8888

ENTRYPOINT ["./hivebox"]


run:
	@go run cmd/hivebox/main.go

build:
	@CGO_ENABLED=0 GOOS=linux go build -o hivebox cmd/hivebox/main.go

.PHONY: run build
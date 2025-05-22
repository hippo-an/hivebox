APP_NAME = 'hivebox'

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

## ==================== Setup ====================

.PHONY: install
install: go-install-air go-install-templ install-tailwindcss ## Install all dependencies

.PHONY: go-install-air
go-install-air:  ## Install air using 'go install'
	go install github.com/air-verse/air@latest

.PHONY: go-install-templ
go-install-templ: ## Install templ using 'go install'
	go install github.com/a-h/templ/cmd/templ@latest

.PHONY: install-tailwindcss
install-tailwindcss: ## Install tailwindcss binary
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-macos-arm64
	chmod +x tailwindcss-macos-arm64
	mv tailwindcss-macos-arm64 twc
	mkdir -p static/css
	echo '@import "tailwindcss";' > static/css/custom.css


## ==================== Build ====================

.PHONY: tailwind-build
tailwind-build: ## Compile Tailwind CSS
	./twc -i ./static/css/custom.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate: ## Generate Go code from templ files
	templ generate -path ./templates

.PHONY: build
build: tailwind-build templ-generate ## Build app and assets
	go build -o ./tmp/$(APP_NAME) ./cmd/hivebox/main.go


## ==================== Live Development ====================

.PHONY: live
live: ## Live reload with air + tailwind + templ
	@make -j4 live/tailwind live/templ live/static live/server 

live/server:
	air -c .air.toml

live/tailwind:
	./twc -i ./static/css/custom.css -o ./static/css/style.css --minify --watch

live/templ:
	templ generate -path ./templates -watch --proxy=http://localhost:8888 --proxybind=localhost -v

live/static:
	air --build.cmd "templ generate --notify-proxy" \
		--build.bin "true" \
		--build.delay "100" \
		--build.exclude_dir "" \
		--build.include_dir "static" \
		--build.include_ext "js,css" \
		--proxy.proxy_port 9099

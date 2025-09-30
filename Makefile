run:
	go run cmd/main.go

swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swag_init:
	@if ! command -v swag >/dev/null 2>&1; then \
		echo "Installing swag..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g api/router.go -o api/docs --parseVendor; \
	else \
		$(shell go env GOPATH)/bin/swag init -g api/router.go -o api/docs --parseVendor; \
	fi

# Docker commands
docker-build:
	docker build -t kizen-go-service .

docker-run:
	docker run -p 8080:8080 --env-file .env kizen-go-service

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-rebuild:
	docker-compose down && docker-compose build --no-cache && docker-compose up -d

clean:
	docker-compose down -v
	docker system prune -f
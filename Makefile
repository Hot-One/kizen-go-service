run:
	go run cmd/main.go

swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest

swag-init:
	swag init -g cmd/main.go -o api/docs --parseDependency --parseInternal


# Docker commands
docker-build:
	docker build -t kizen-go-service .

docker-run:
	docker run -p 8080:8080 --env-file .env kizen-go-service

# Docker Compose commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-rebuild:
	docker-compose down && docker-compose build --no-cache && docker-compose up -d

# Development commands
dev-up:
	docker-compose -f docker-compose.dev.yml up -d

dev-down:
	docker-compose -f docker-compose.dev.yml down

dev-logs:
	docker-compose -f docker-compose.dev.yml logs -f

# Database commands
db-migrate:
	docker-compose exec kizen-service ./main migrate

db-seed:
	docker-compose exec kizen-service ./main seed

# Cleanup
clean:
	docker-compose down -v
	docker system prune -f
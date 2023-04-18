# Load environment variables from .env file
include dev.env

export

# Run the app
run:
	docker compose -f docker-compose.dev.yaml up -d
	go run cmd/myapp/main.go
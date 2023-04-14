# Load environment variables from .env file
include .env

export

# Run the app
run:
	docker compose up -d
	go run cmd/myapp/main.go
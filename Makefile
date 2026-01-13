.PHONY: db api web up down

# Start the db container, listening on port 5432
up:
	docker compose up -d

# Stop and remove the db container
down:
	docker compose down -v

# View the db logs
db:
	docker compose logs -f db

# Run the Go API server
api:
	cd api && go run .

# Run the React web app
web:
	cd web && yarn run dev

# Install dependencies for the React web app
install:
	cd web && yarn install
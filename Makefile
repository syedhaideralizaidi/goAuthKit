# Load environment variables from .env file
include .env

# Directory for migration files
MIGRATION_DIR=internal/database/migrations

# Migration extension
MIGRATION_EXT=sql

# Docker PostgreSQL image version
POSTGRES_IMAGE=postgres:13

# Ensure environment variables from .env are used
POSTGRES_USER := $(POSTGRES_USER)
POSTGRES_PASSWORD := $(POSTGRES_PASSWORD)
POSTGRES_CONTAINER_NAME := $(POSTGRES_CONTAINER_NAME)
DATABASE_NAME := $(DATABASE_NAME)
DATABASE_URL := $(DATABASE_URL)

# Targets
.PHONY: create_migration postgres_up postgres_down createdb dropdb migrate_up migrate_down generate_sqlc access_db run

# Create a new migration file
create_migration:
	migrate create -ext=$(MIGRATION_EXT) -dir=$(MIGRATION_DIR) -seq init

# Start PostgreSQL container
postgres_up:
	docker run --name $(POSTGRES_CONTAINER_NAME) -p 5432:5432 -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -d $(POSTGRES_IMAGE)

# Stop and remove PostgreSQL container
postgres_down:
	docker stop $(POSTGRES_CONTAINER_NAME)
	docker rm $(POSTGRES_CONTAINER_NAME)

# Create the database
createdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(DATABASE_NAME)

# Drop the database
dropdb:
	docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb $(DATABASE_NAME)

# Apply all migrations
migrate_up:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" -verbose up

# Rollback all migrations
migrate_down:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE_URL)" -verbose down

# Generate SQL code using sqlc
generate_sqlc:
	sqlc generate

# Access the PostgreSQL database
access_db:
	docker exec -it $(POSTGRES_CONTAINER_NAME) psql -U $(POSTGRES_USER) -d $(DATABASE_NAME)

# Run the main Go application
run:
	go run cmd/authKit/main.go
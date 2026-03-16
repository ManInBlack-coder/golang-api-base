.PHONY: seed migrate up down build


up:
	cd .. && docker-compose up -d

down:
	cd .. && docker-compose down

build:
	cd .. && docker-compose up -d --build

seed:
	go run cmd/seed/main.go

migrate:
	# Run migrations (can be extended with migration tool)
	psql -h localhost -U postgres -d go_base_api -f migrations/001_create_users_table.sql
	psql -h localhost -U postgres -d go_base_api -f migrations/002_create_indexes.sql

logs:
	cd .. && docker-compose logs -f api

ps:
	cd .. && docker-compose ps

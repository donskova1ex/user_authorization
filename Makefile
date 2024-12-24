include scripts/*.mk

DEV_COMPOSE_ARGS=--env-file .env.dev -f docker-compose.dev.yaml
DEV_COMPOSE_ENV=docker compose $(DEV_COMPOSE_ARGS)
DEV_COMPOSE=docker compose --profile api $(DEV_COMPOSE_ARGS)

dev-build:
	$(DEV_COMPOSE) build

dev-up: api_docker_build dev-build
	$(DEV_COMPOSE) up -d

dev-up-env: dev-build
	$(DEV_COMPOSE_ENV) up -d

dev-down:
	$(DEV_COMPOSE) down

dev-logs:
	docker compose -f docker-compose.dev.yaml logs

dev-migrate-up:
	docker-compose -f docker-compose.dev.yaml up -d migrations-up

dev-migrate-down:
	docker compose --profile migrations-down -f docker-compose.dev.yaml up -d migrations-down

dev-api-run:
	go run cmd/api/api.go

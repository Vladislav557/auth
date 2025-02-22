DB ?= postgresql://dev:dev@172.1.10.1:54323/auth?sslmode=disable
MIGRATION_NAME ?=

jwt-generate:
	./generator.sh

create_network:
	docker network create --driver bridge --subnet 172.1.10.0/24 --gateway 172.1.10.1 docker_net

build:
	docker compose up -d --build

setup-migrator:
	go install github.com/pressly/goose/v3/cmd/goose@latest

create-migration:
	goose create $(MIGRATION_NAME) sql

migrate-up:
	goose -dir migrations postgres "$(DB)" up

migrate-down:
	goose -dir migrations postgres "$(DB)" down

run:
	go run github.com/Vladislav557/auth/cmd/auth

smtp:
	docker run -d -p 1025:1025 -p 8025:8025 mailhog/mailhog
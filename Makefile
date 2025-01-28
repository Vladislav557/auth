DB ?= postgresql://dev:dev@172.1.10.1:54323/dev?sslmode=disable

jwt-generate:
	./generator.sh

create_network:
	docker network create --driver bridge --subnet 172.1.10.0/24 --gateway 172.1.10.1 docker_net

build:
	docker compose up -d --build

setup-migrator:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrate-up:
	goose -dir migrations postgres "$(DB)" up

migrate-down:
	goose -dir migrations postgres "$(DB)" down

run:
	go build github.com/Vladislav557/auth/cmd/auth
jwt-generate:
	./generator.sh

create_network:
	docker network create --driver bridge --subnet 172.1.10.0/24 --gateway 172.1.10.1 docker_net

build:
	docker compose up -d --build
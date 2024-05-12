.PHONY: build
build:
	docker build --progress=plain -t service:latest .

.PHONY: up
up:
	docker-compose up --force-recreate -d

.PHONY: down
down:
	docker-compose down --volumes
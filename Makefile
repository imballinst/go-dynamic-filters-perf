.PHONY: build
build:
	docker build --progress=plain -t service:latest .

.PHONY: up
up:
	docker-compose up

.PHONY: down
down:
	docker-compose down --volumes
	
.PHONY: dev
dev:
	docker-compose -f docker-compose.dev.yaml up
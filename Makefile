DOCKER_COMPOSE_CMD=docker compose

install:
	$(DOCKER_COMPOSE_CMD) run cmd go mod tidy

run:
	@$(DOCKER_COMPOSE_CMD) run cmd go run main.go

test:
	@$(DOCKER_COMPOSE_CMD) run cmd go test -v ./...

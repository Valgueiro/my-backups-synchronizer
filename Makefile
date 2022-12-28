RCLONE_CONFIG_PATH=/home/mvalgueiro/.config/rclone/rclone.conf
DOCKER_COMPOSE_CMD=RCLONE_CONFIG_PATH=$(RCLONE_CONFIG_PATH) docker compose

build:
	$(DOCKER_COMPOSE_CMD) build

install: build
	$(DOCKER_COMPOSE_CMD) run cmd go mod tidy

run:
	@$(DOCKER_COMPOSE_CMD) run cmd go run main.go

test:
	@$(DOCKER_COMPOSE_CMD) run cmd go test -v ./...

bash:
	$(DOCKER_COMPOSE_CMD) run cmd bash

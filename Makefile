RCLONE_CONFIG_PATH=/home/mvalgueiro/.config/rclone/
ARCHIVE_PATH=/media/mvalgueiro/Arquivos1/Backups


ENV=RCLONE_CONFIG_PATH=$(RCLONE_CONFIG_PATH) ARCHIVE_PATH=$(ARCHIVE_PATH)
DOCKER_COMPOSE_CMD=$(ENV) docker compose


up:
	$(DOCKER_COMPOSE_CMD) up -d --build

down:
	$(DOCKER_COMPOSE_CMD) down

ps:
	$(DOCKER_COMPOSE_CMD) ps

build:
	$(DOCKER_COMPOSE_CMD) build --no-cache

install:
	$(DOCKER_COMPOSE_CMD) exec cmd go mod tidy

run:
	@$(DOCKER_COMPOSE_CMD) exec cmd go run main.go

test:
	$(DOCKER_COMPOSE_CMD) exec cmd go test -v ./...

bash:
	$(DOCKER_COMPOSE_CMD) exec cmd bash

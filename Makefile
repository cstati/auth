CONFIGS_PATH=./configs

.PHONY: build
build:
	go build -o ./bin/auth ./cmd/auth/main.go

.PHONY: run
run:
	DOTENV_FILE=$(CONFIGS_PATH)/dev/.env go run ./cmd/auth/main.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: imports
imports:
	gci write  .
	 goimports -w .

.PHONY: docker-build
docker-build:
	docker build -t auth:latest -f ./build/Dockerfile .

.PHONY: docker-run
docker-run:
	docker compose -f deployments/dev/docker-compose.yaml up --build

.PHONY: docker-run
docker-run-background:
	docker compose -f deployments/dev/docker-compose.yaml up --build -d

.PHONY: start-infra
start-infra:
	docker compose -f deployments/dev/docker-compose.yaml up --build -d db-migrator

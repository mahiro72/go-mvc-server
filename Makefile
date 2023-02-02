.PHONY: up
up: check-env ## docker compose環境を立ち上げる
	docker-compose up -d

.PHONY: down
down: ## docker compose環境を閉じる
	docker-compose down

.PHONY: log
log: ## docker compose環境のログを標示する
	docker-compose logs -f

.PHONY: check-env
check-env:
	./scripts/check-env.sh

.PHONY: help
help: ## helpを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'


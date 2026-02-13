.PHONY: help build run clean deps fmt vet test

# Переменные
APP_NAME=platform
BUILD_DIR=bin
CMD_DIR=cmd/platform
MAIN_FILE=$(CMD_DIR)/main.go

help: ## Показать справку
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

deps: ## Установить зависимости
	@echo "Загрузка зависимостей..."
	go mod download
	go mod tidy

build: ## Собрать приложение
	@echo "Сборка приложения..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)
	@echo "✓ Приложение собрано: $(BUILD_DIR)/$(APP_NAME)"

run: ## Запустить приложение
	@echo "Запуск приложения..."
	go run $(MAIN_FILE)

run-build: build ## Собрать и запустить
	@echo "Запуск собранного приложения..."
	./$(BUILD_DIR)/$(APP_NAME)

fmt: ## Форматировать код
	@echo "Форматирование кода..."
	go fmt ./...

vet: ## Проверить код
	@echo "Проверка кода..."
	go vet ./...

clean: ## Удалить бинарные файлы
	@echo "Очистка..."
	rm -rf $(BUILD_DIR)
	@echo "✓ Очистка завершена"

test: ## Запустить тесты
	@echo "Запуск тестов..."
	go test -v ./...

setup: ## Начальная настройка проекта
	@echo "Настройка проекта..."
	@if [ ! -f .env ]; then cp .env.example .env; echo "✓ Создан файл .env"; fi
	@$(MAKE) deps
	@echo "✓ Проект настроен"

docker-postgres: ## Запустить PostgreSQL в Docker
	@echo "Запуск PostgreSQL в Docker..."
	docker run --name postgres-dev \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=semen_db \
		-p 5432:5432 \
		-d postgres:15-alpine
	@echo "✓ PostgreSQL запущен на порту 5432"

docker-stop: ## Остановить Docker контейнеры
	@echo "Остановка контейнеров..."
	docker stop postgres-dev || true
	docker rm postgres-dev || true
	@echo "✓ Контейнеры остановлены"

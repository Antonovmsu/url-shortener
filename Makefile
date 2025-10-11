.PHONY: run build test clean

APP_NAME=url-shortener

# Запуск приложения
run:
	go run cmd/$(APP_NAME)/main.go

# Сборка приложения
build:
	go build -o bin/$(APP_NAME) cmd/$(APP_NAME)/main.go

# Запуск тестов
test:
	go test ./...

# Очистка билдов
clean:
	rm -rf bin/

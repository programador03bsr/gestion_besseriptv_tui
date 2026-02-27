# Variables
BINARY_NAME=canales_cli
MAIN_PATH=cmd/canales_cli/main.go

.PHONY: all build run clean lint fmt test hooks

all: lint build

build:
	@echo "==> Construyendo el binario..."
	@mkdir -p bin
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "==> Listo: bin/$(BINARY_NAME)"

run: build
# 	@echo "==> Ejecutando TUI..."
	@./bin/$(BINARY_NAME)

clean:
	@echo "==> Limpiando..."
	go clean
	rm -rf bin/

lint:
	@echo "==> Ejecutando golangci-lint..."
	golangci-lint run

fmt:
	@echo "==> Formateando código con gofumpt..."
	gofumpt -l -w .

test:
	@echo "==> Ejecutando pruebas..."
	go test ./... -v

hooks:
	@echo "==> Configurando Git Hooks locales..."
	chmod +x .githooks/pre-commit .githooks/commit-msg
	git config core.hooksPath .githooks
	@echo "==> Hooks activados."
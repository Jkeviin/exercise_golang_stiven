.PHONY: help run dev test clean deps

help:
	@echo "📋 Comandos disponibles:"
	@echo "  make deps     - Instala dependencias"
	@echo "  make dev      - 🔥 Ejecuta servidor con hot reload (Air)"
	@echo "  make run      - Ejecuta el servidor (sin hot reload)"
	@echo "  make test     - Ejecuta los tests"
	@echo "  make clean    - Limpia archivos compilados"
	@echo "  make build    - Compila el ejecutable"

deps:
	@echo "📦 Instalando dependencias..."
	go mod tidy
	go mod download

dev:
	@echo "🔥 Iniciando servidor con hot reload (Air)..."
	@echo "💡 Los cambios se recargarán automáticamente"
	@which air > /dev/null || (echo "❌ Air no está instalado. Ejecuta: go install github.com/air-verse/air@latest" && exit 1)
	air

run:
	@echo "🚀 Iniciando servidor..."
	go run cmd/app/main.go

test:
	@echo "✅ Ejecutando tests..."
	go test ./internal/usecase/... -v

clean:
	@echo "🧹 Limpiando..."
	rm -f ejercicio-api
	go clean

build:
	@echo "🔨 Compilando..."
	go build -o ejercicio-api cmd/app/main.go
	@echo "✅ Ejecutable creado: ./ejercicio-api"


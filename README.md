# 🚀 Ejercicio API - Clean Architecture en Golang

Proyecto educativo para aprender Clean Architecture y DDD desde cero.

## ⚡ Inicio Rápido

### Mac/Linux:

```bash
# Instalar dependencias
go mod tidy

# Ejecutar servidor
go run cmd/app/main.go

# Probar endpoints
curl http://localhost:8080/status
curl http://localhost:8080/ping
curl http://localhost:8080/users/1
```

### Windows:

```cmd
REM Instalar dependencias
scripts\deps.bat

REM Ejecutar servidor
scripts\run.bat

REM O ejecutar con hot reload
scripts\dev.bat
```

> **💡 Nota para Windows**: `make` no funciona por defecto. Usa los scripts `.bat` en la carpeta `scripts/`. Ver [Guía Windows](docs/WINDOWS.md).

## 🔥 Desarrollo con Hot Reload

Para no tener que reiniciar el servidor en cada cambio, usa **Air**:

```bash
# 1. Instalar Air (solo una vez)
go install github.com/air-verse/air@latest

# 2. Ejecutar con hot reload
air
# O usando Make:
make dev
```

**Con Air**:
- ✅ Reinicio automático al guardar cambios
- ✅ No necesitas detener/iniciar el servidor
- ✅ Compilación rápida
- ✅ Logs en colores

## 🧪 Tests

```bash
go test ./...  -v
```

## 📚 Documentación

- **[ARQUITECTURA.md](ARQUITECTURA.md)** - Explicación detallada de cómo está construido el proyecto, responsabilidad de cada carpeta según DDD y Clean Architecture
- **[WINDOWS.md](docs/WINDOWS.md)** - 🪟 Guía completa para usar el proyecto en Windows (scripts .bat, make alternativo)
- **[HOT_RELOAD.md](docs/HOT_RELOAD.md)** - 🔥 Guía de hot reload con Air (recarga automática sin reiniciar)
- **[WORKSHOP.md](docs/WORKSHOP.md)** - Ejercicios prácticos paso a paso para aprender
- **[README_POSTMAN.md](README_POSTMAN.md)** - Guía para usar la colección de Postman

## 📁 Estructura (Clean Architecture + DDD)

```
ejercicio-api/
├── cmd/app/                              # Punto de entrada
├── internal/
│   ├── domain/                           # CAPA DE DOMINIO
│   │   ├── user/
│   │   │   ├── user.go                   # Entidad User
│   │   │   └── repository.go             # Interface del repositorio
│   │   ├── status/
│   │   │   └── status.go                 # Entidad Status
│   │   └── ping/
│   │       └── ping.go                   # Entidad Ping
│   │
│   ├── usecase/                          # CAPA DE APLICACIÓN
│   │   ├── user/
│   │   │   └── get_user.go               # Caso de uso: Obtener usuario
│   │   ├── status/
│   │   │   └── get_status.go             # Caso de uso: Obtener status
│   │   └── ping/
│   │       └── ping.go                   # Caso de uso: Ping
│   │
│   ├── adapter/                          # CAPA DE ADAPTADORES
│   │   ├── http/handler/                 # Handlers HTTP
│   │   │   ├── user_handler.go
│   │   │   ├── status_handler.go
│   │   │   └── ping_handler.go
│   │   └── repository/
│   │       └── user_api_repository.go    # Implementación del repositorio
│   │
│   ├── infrastructure/                   # CAPA DE INFRAESTRUCTURA
│   │   └── http/
│   │       └── router.go                 # Router y servidor HTTP
│   │
│   └── config/
│       └── config.go                     # Configuración
│
└── test/                                 # Tests organizados por usecase
    └── usecase/
        ├── user/
        ├── status/
        └── ping/
```

## 🎯 Endpoints

| Ruta | Descripción |
|------|-------------|
| `GET /status` | Estado del servidor con uptime |
| `GET /ping` | Health check |
| `GET /users/{id}` | Usuario por ID (API externa) |

## 🏛️ Principios Aplicados

### Clean Architecture
- ✅ **Independencia de frameworks**: El dominio no depende de chi, http, etc.
- ✅ **Testeable**: Los casos de uso se prueban sin necesidad de servidor
- ✅ **Independiente de UI**: Los handlers son intercambiables
- ✅ **Independiente de BD**: El repositorio es una interfaz

### DDD (Domain-Driven Design)
- ✅ **Entidades en domain**: `User`, `Status`, `Ping`
- ✅ **Repositorios como interfaces**: `user.Repository`
- ✅ **Casos de uso**: Lógica de aplicación separada
- ✅ **Adaptadores**: Implementaciones concretas fuera del dominio

## 🔄 Flujo de Datos

```
HTTP Request
    ↓
Handler (adapter/http/handler)
    ↓
Usecase (usecase/)
    ↓
Repository Interface (domain/)
    ↓
Repository Implementation (adapter/repository)
    ↓
External API
    ↓
Domain Entity
    ↓
Response
```

## 🛠 Tecnologías

- **Go 1.21+**
- **chi router** - Router HTTP moderno
- **Clean Architecture** - Separación en capas
- **DDD** - Domain-Driven Design

## 📖 Aprendizaje

Ve a [docs/WORKSHOP.md](docs/WORKSHOP.md) para ejercicios prácticos paso a paso.

## ⚙️ Configuración

Variables de entorno:

```bash
SERVER_PORT=8080
EXTERNAL_API_URL=https://jsonplaceholder.typicode.com
```

## 🔧 Comandos

### Mac/Linux (con Make):

```bash
make deps     # Instalar dependencias
make dev      # 🔥 Ejecutar con hot reload (recomendado para desarrollo)
make run      # Ejecutar servidor (sin hot reload)
make test     # Ejecutar tests
make build    # Compilar ejecutable
```

### Windows (scripts .bat):

```cmd
scripts\deps.bat      # Instalar dependencias
scripts\dev.bat       # 🔥 Ejecutar con hot reload
scripts\run.bat       # Ejecutar servidor (sin hot reload)
scripts\test.bat      # Ejecutar tests
scripts\build.bat     # Compilar ejecutable
scripts\help.bat      # Ver ayuda
```

## 🌟 Características Destacadas

- ✅ Arquitectura hexagonal completa
- ✅ Inyección de dependencias
- ✅ Interfaces para testabilidad
- ✅ Tests aislados (mocks)
- ✅ Separación clara de responsabilidades
- ✅ Código escalable y mantenible

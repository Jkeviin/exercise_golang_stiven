# 🚀 Ejercicio API - Clean Architecture en Go

Proyecto educativo para aprender Clean Architecture y DDD desde cero con 38 ejercicios prácticos.

---

## ⚡ Inicio Rápido

```bash
# 1. Instalar dependencias
go mod tidy

# 2. Ejecutar servidor
go run cmd/app/main.go

# 3. Probar endpoints
curl http://localhost:8080/          # Bienvenida
curl http://localhost:8080/status    # Estado del servidor
curl http://localhost:8080/ping      # Health check
curl http://localhost:8080/users     # Lista de usuarios
curl http://localhost:8080/users/1   # Usuario específico
```

**Si funciona todo** → ¡Adelante! 🎉

---

## 📚 Documentación

| Archivo | Descripción |
|---------|-------------|
| **[docs/TALLER.md](docs/TALLER.md)** | 🎓 **38 ejercicios prácticos** - Empieza aquí |
| **[ARQUITECTURA.md](ARQUITECTURA.md)** | 🏛️ Explicación de Clean Architecture + DDD |
| **[docs/ENDPOINTS.md](docs/ENDPOINTS.md)** | 📍 Guía de endpoints disponibles |

---

## 🎯 Endpoints Disponibles

| Ruta | Método | Descripción |
|------|--------|-------------|
| `/` | GET | Mensaje de bienvenida |
| `/status` | GET | Estado del servidor con métricas |
| `/ping` | GET | Health check |
| `/users` | GET | Lista todos los usuarios |
| `/users/{id}` | GET | Usuario por ID (1-10) |

---

## 📁 Estructura del Proyecto

```
ejercicio-api/
├── cmd/app/main.go         # Punto de entrada
├── internal/
│   ├── domain/             # Entidades y reglas de negocio
│   ├── usecase/            # Casos de uso (lógica de aplicación)
│   ├── adapter/            # Adaptadores (handlers, repositorios)
│   ├── infrastructure/     # Infraestructura (router, servidor)
│   └── config/             # Configuración
├── docs/
│   ├── TALLER.md          # 38 ejercicios prácticos
│   └── ENDPOINTS.md       # Documentación de endpoints
└── ARQUITECTURA.md        # Guía de arquitectura
```

---

## 🏛️ Arquitectura

Este proyecto implementa:

- ✅ **Clean Architecture** - Separación en capas independientes
- ✅ **Domain-Driven Design (DDD)** - Código organizado por dominio
- ✅ **Inyección de Dependencias** - Componentes desacoplados
- ✅ **Arquitectura Hexagonal** - Fácil de testear y mantener

**Flujo de datos:**
```
HTTP Request → Handler → UseCase → Repository → API Externa → Response
```

---

## 🎓 Aprendizaje

### ¿Quieres aprender haciendo?

Ve a **[docs/TALLER.md](docs/TALLER.md)** para:
- 38 ejercicios progresivos
- Desde básico hasta avanzado
- Sin código para copiar (aprende de verdad)

**No copies y pegues - piensa, intenta, debuggea, aprende.**

---

## 🔥 Desarrollo con Hot Reload (Opcional)

Para no reiniciar el servidor en cada cambio:

```bash
# 1. Instalar Air (solo una vez)
go install github.com/air-verse/air@latest

# 2. Ejecutar con hot reload
air
```

---

## 🧪 Tests

```bash
go test ./... -v
```

---

## ⚙️ Configuración

Variables de entorno:

```bash
SERVER_PORT=8080
EXTERNAL_API_URL=https://jsonplaceholder.typicode.com
```

---

## 🛠 Tecnologías

- **Go 1.21+**
- **chi router** - Router HTTP minimalista
- **Clean Architecture** - Separación en capas
- **DDD** - Domain-Driven Design

---

## 🌟 Características

- ✅ Arquitectura profesional y escalable
- ✅ Código limpio y bien organizado
- ✅ Fácil de entender y mantener
- ✅ Preparado para tests
- ✅ 38 ejercicios prácticos para aprender

---

## 📖 Recursos de Aprendizaje

1. **[docs/TALLER.md](docs/TALLER.md)** - Empieza aquí con los ejercicios
2. **[ARQUITECTURA.md](ARQUITECTURA.md)** - Entiende cómo está construido
3. **[docs/ENDPOINTS.md](docs/ENDPOINTS.md)** - Explora los endpoints

---

## 🚀 Siguiente Paso

```bash
# Empieza el taller
cat docs/TALLER.md
```

¡Buena suerte! 💪

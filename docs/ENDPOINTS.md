# 📍 Guía de Endpoints

Este documento explica el propósito y las diferencias entre cada endpoint de la API.

---

## 🔍 Endpoints de Monitoreo y Salud

### `GET /ping` - Health Check Rápido

**Propósito**: Verificación mínima de que el servidor está respondiendo

**Cuándo usarlo**:
- Load balancers que necesitan verificar si el servidor está vivo
- Monitoreo automático cada pocos segundos
- Scripts de health check simples

**Respuesta**:
```json
{
  "message": "API funcionando correctamente"
}
```

**Características**:
- ✅ Ultra rápido (sin lógica compleja)
- ✅ Sin dependencias externas
- ✅ Solo verifica que el servidor responde

---

### `GET /status` - Estado Detallado del Servidor

**Propósito**: Monitoreo completo con métricas del servidor

**Cuándo usarlo**:
- Dashboards de monitoreo
- Debugging de problemas
- Verificar uptime y versión

**Respuesta**:
```json
{
  "message": "La aplicación está funcionando correctamente",
  "version": "1.1.0",
  "uptime": 3600,
  "environment": "development",
  "timestamp": "2025-12-15T10:30:00Z"
}
```

**Características**:
- ✅ Métricas detalladas (uptime, versión)
- ✅ Información del ambiente
- ✅ Timestamp actual
- ✅ Útil para debugging

---

### `GET /` - Bienvenida y Documentación

**Propósito**: Punto de entrada informativo para desarrolladores

**Cuándo usarlo**:
- Primera vez que alguien accede a la API
- Descubrir qué endpoints están disponibles
- Verificar versión de la API

**Respuesta**:
```json
{
  "message": "Bienvenido a la API de Ejercicio",
  "version": "1.1.0",
  "endpoints": [
    "/",
    "/status",
    "/ping",
    "/users",
    "/users/{id}"
  ]
}
```

**Características**:
- ✅ Lista de todos los endpoints disponibles
- ✅ Mensaje amigable para humanos
- ✅ Útil como "auto-documentación"

---

## 🆚 Comparación: ¿Cuál usar?

| Aspecto | `/ping` | `/status` | `/` |
|---------|---------|-----------|-----|
| **Propósito** | Health check | Monitoreo | Documentación |
| **Audiencia** | Máquinas | Ops/DevOps | Desarrolladores |
| **Rapidez** | ⚡⚡⚡ | ⚡⚡ | ⚡⚡ |
| **Información** | Mínima | Detallada | Navegacional |
| **Uso típico** | Load balancer | Dashboard | Browser |

---

## 👥 Endpoints de Negocio

### `GET /users` - Listar Usuarios

**Propósito**: Obtener todos los usuarios disponibles

**Respuesta**:
```json
[
  {
    "id": 1,
    "name": "Leanne Graham",
    "email": "Sincere@april.biz",
    "username": "Bret"
  },
  ...
]
```

---

### `GET /users/{id}` - Obtener Usuario por ID

**Propósito**: Obtener un usuario específico

**Parámetros**:
- `id` (path): ID del usuario (1-10)

**Ejemplo**: `GET /users/1`

**Validaciones**:
- ❌ ID debe ser un número
- ❌ ID debe ser mayor que 0
- ❌ ID debe estar entre 1 y 10

**Respuestas**:
- 200: Usuario encontrado
- 400: ID inválido
- 404: Usuario no encontrado

---

## 🎯 Flujo Recomendado para Nuevos Desarrolladores

1. **Primer contacto**: Visita `GET /` para ver qué está disponible
2. **Verificar salud**: Llama `GET /status` para confirmar versión y estado
3. **Probar datos**: Usa `GET /users` o `GET /users/1` para ver datos reales
4. **Integrar**: Configura `GET /ping` en tu load balancer

---

## 📝 Notas de Diseño

### ¿Por qué tres endpoints similares?

Aunque `/ping`, `/status` y `/` pueden parecer redundantes, cada uno sirve un propósito diferente:

- **`/ping`**: Diseñado para ser llamado miles de veces por minuto sin impacto
- **`/status`**: Proporciona métricas útiles para operaciones
- **`/`**: Ayuda a desarrolladores humanos a explorar la API

En un sistema real de producción, los tres son estándares de la industria y se usan simultáneamente por diferentes sistemas.

---

## 🔮 Futuros Endpoints (Workshop)

Durante el workshop, aprenderás a crear:

- `GET /users/search?email=...` - Buscar por email
- `GET /users/{id}/email` - Solo obtener email
- Y muchos más...

Cada ejercicio del WORKSHOP.md te enseñará a crear endpoints siguiendo el mismo patrón limpio y escalable.


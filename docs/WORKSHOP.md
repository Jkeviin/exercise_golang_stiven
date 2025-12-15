# 🎓 TALLER PRÁCTICO - APIs REST con Clean Architecture

**38 ejercicios progresivos** para dominar desarrollo de APIs en Go.

---

## ⚡ METODOLOGÍA

### ❌ NO copies y pegues código
### ✅ LEE, PIENSA, INTENTA, DEBUGGEA, APRENDE

**Cada ejercicio tiene:**
1. **Descripción del problema** - ¿Qué necesita el cliente?
2. **Pistas mínimas** - Orientación sin revelar la solución
3. **Referencias** - Qué archivos estudiar primero

**Tiempo sugerido por ejercicio:**
- Básicos (1-8): 10-20 minutos
- Intermedios (9-30): 30-60 minutos
- Avanzados (31-38): 60-90 minutos

---

## 🚀 ANTES DE EMPEZAR

### Verifica que todo funciona

```bash
# 1. Inicia el servidor
go run cmd/app/main.go

# 2. En otra terminal, prueba:
curl http://localhost:8080/status    # Estado del servidor
curl http://localhost:8080/ping      # Health check
curl http://localhost:8080/users/1   # Usuario por ID
```

Si todo funciona → ¡Adelante! 🎉

---

## 📚 CONCEPTOS CLAVE

### Flujo de Desarrollo (en este orden)

```
1. RUTA (router.go)
   ↓ "Defino QUÉ endpoint quiero: GET /users"
   
2. HANDLER (handler/)
   ↓ "Manejo la petición HTTP: extraigo parámetros, valido formato"
   
3. CASO DE USO (usecase/)
   ↓ "Aplico lógica de negocio y validaciones"
   
4. REPOSITORIO (repository/)
   ↓ "Obtengo/guardo datos (API externa, BD)"
   
5. MAIN (cmd/app/main.go)
   ↓ "Conecto todo con inyección de dependencias"
```

### Estructura del Proyecto

```
ejercicio-api/
├── cmd/app/main.go              # Punto de entrada
├── internal/
│   ├── domain/                  # Entidades (User, Product, Status)
│   ├── usecase/                 # Lógica de negocio
│   ├── handler/                 # Adaptadores HTTP
│   ├── repository/              # Acceso a datos
│   └── router/                  # Configuración de rutas
```

**Lee también:** `ARQUITECTURA.md` para entender arquitectura hexagonal completa.

---

## 🟢 NIVEL BÁSICO (Ejercicios 1-8)

Modificar código existente, agregar validaciones simples.

---

### EJERCICIO 1 - Cambiar el mensaje del ping

#### 📋 Historia de Usuario
```
Como líder técnico
Quiero que el endpoint /ping devuelva un mensaje más descriptivo
Para que sea más claro para los nuevos desarrolladores qué significa la respuesta
```

#### ✅ Criterios de Aceptación

**CA1:** El endpoint `/ping` debe devolver mensaje descriptivo
- Antes: `{"message": "pong"}`
- Ahora: `{"message": "API funcionando correctamente"}`

**CA2:** La estructura JSON debe mantenerse igual
- Solo cambia el texto del mensaje
- El campo sigue llamándose `message`

#### 🧪 Escenarios de Prueba

**Escenario 1: Verificar nuevo mensaje**
```bash
curl http://localhost:8080/ping
# Esperado: {"message":"API funcionando correctamente"}
```

#### 💡 Pistas Técnicas
- Archivo a modificar: `internal/usecase/ping/ping.go`
- Busca la línea que contiene `"pong"`
- Reemplaza el texto por `"API funcionando correctamente"`
- Reinicia el servidor para ver el cambio

---

### EJERCICIO 2 - Actualizar versión de la API

#### 📋 Historia de Usuario
```
Como equipo de producto
Queremos actualizar el número de versión en el endpoint /status
Para reflejar las mejoras implementadas en el sistema
```

#### ✅ Criterios de Aceptación

**CA1:** El campo `version` debe mostrar la nueva versión
- Antes: `"version": "1.0.0"`
- Ahora: `"version": "1.1.0"`

**CA2:** Solo debe cambiar el número de versión
- Los demás campos del `/status` permanecen iguales

#### 🧪 Escenarios de Prueba

**Escenario 1: Verificar nueva versión**
```bash
curl http://localhost:8080/status
# El campo "version" debe mostrar "1.1.0"
```

#### 💡 Pistas Técnicas
- Archivo: `internal/usecase/status/get_status.go`
- Busca: `Version: "1.0.0"`
- Cámbialo a: `Version: "1.1.0"`

---

### EJERCICIO 3 - Validar IDs negativos

#### 📋 Historia de Usuario
```
Como usuario de la API
Quiero recibir un error claro cuando solicito un usuario con ID negativo
Para entender rápidamente que mi petición es inválida
```

#### ✅ Criterios de Aceptación

**CA1:** Rechazar IDs menores o iguales a cero
- `/users/0` → Error
- `/users/-1` → Error
- `/users/-999` → Error

**CA2:** Devolver mensaje de error descriptivo
- HTTP 400 Bad Request
- Mensaje: "el ID debe ser mayor que 0" (o similar)

**CA3:** IDs positivos deben seguir funcionando
- `/users/1` → OK (devuelve usuario)
- `/users/5` → OK (devuelve usuario)

#### 🧪 Escenarios de Prueba

**Escenario 1: ID negativo**
```bash
curl http://localhost:8080/users/-5
# Esperado: Error con mensaje claro
```

**Escenario 2: ID cero**
```bash
curl http://localhost:8080/users/0
# Esperado: Error con mensaje claro
```

**Escenario 3: ID positivo (debe seguir funcionando)**
```bash
curl http://localhost:8080/users/1
# Esperado: Información del usuario
```

#### 💡 Pistas Técnicas
- Archivo: `internal/usecase/user/get_user.go`
- Agregar validación al inicio del método `Execute()`
- Usar: `if id <= 0 { return nil, errors.New("el ID debe ser mayor que 0") }`

---

### EJERCICIO 4 - Limitar IDs al rango disponible

#### 📋 Historia de Usuario
```
Como administrador de sistemas
Quiero que la API rechace IDs fuera del rango válido (1-10)
Para evitar llamadas innecesarias a la API externa que sabemos que fallarán
```

#### ✅ Criterios de Aceptación

**CA1:** Rechazar IDs mayores a 10
- La API externa `jsonplaceholder` solo tiene usuarios del 1 al 10
- `/users/11` → Error
- `/users/999` → Error

**CA2:** Mensaje de error específico
- HTTP 400 Bad Request
- Mensaje: "el ID debe estar entre 1 y 10" (o similar)

**CA3:** Rango válido funciona correctamente
- `/users/1` a `/users/10` → OK

#### 🧪 Escenarios de Prueba

**Escenario 1: ID fuera de rango superior**
```bash
curl http://localhost:8080/users/11
# Esperado: Error "ID debe estar entre 1 y 10"
```

**Escenario 2: ID muy grande**
```bash
curl http://localhost:8080/users/999
# Esperado: Error "ID debe estar entre 1 y 10"
```

**Escenario 3: Límite superior válido**
```bash
curl http://localhost:8080/users/10
# Esperado: Usuario devuelto correctamente
```

**Escenario 4: Combinar con ejercicio 3**
```bash
curl http://localhost:8080/users/0
# Esperado: Error de ID <= 0 (validación del ejercicio 3)
```

#### 💡 Pistas Técnicas
- Mismo archivo que ejercicio 3: `internal/usecase/user/get_user.go`
- Agregar DESPUÉS de la validación `id <= 0`
- Código: `if id > 10 { return nil, errors.New("el ID debe estar entre 1 y 10") }`

---

### EJERCICIO 5 - Validar IDs no numéricos

#### 📋 Historia de Usuario
```
Como desarrollador frontend
Quiero recibir un error claro cuando envío un ID que no es número
Para poder validar mejor los datos de entrada en mi aplicación
```

#### ✅ Criterios de Aceptación

**CA1:** Rechazar IDs que no sean números válidos
- `/users/abc` → Error
- `/users/hola` → Error
- `/users/@#$` → Error

**CA2:** Mensaje de error apropiado
- HTTP 400 Bad Request
- Mensaje: "ID inválido" o "El ID debe ser un número"

**CA3:** IDs numéricos válidos siguen funcionando
- `/users/1` → OK
- `/users/10` → OK

#### 🧪 Escenarios de Prueba

**Escenario 1: ID con letras**
```bash
curl http://localhost:8080/users/abc
# Esperado: HTTP 400 - "ID inválido"
```

**Escenario 2: ID con caracteres especiales**
```bash
curl http://localhost:8080/users/@#$
# Esperado: HTTP 400 - "ID inválido"
```

**Escenario 3: ID numérico válido**
```bash
curl http://localhost:8080/users/5
# Esperado: Información del usuario
```

#### 💡 Pistas Técnicas
- Archivo: `internal/handler/user_handler.go`
- Busca donde se convierte el ID con `strconv.Atoi()`
- Verifica que exista: `if err != nil { ... }`
- Si no existe, agrégalo con mensaje claro

---

### EJERCICIO 6 - Indicar ambiente de ejecución

#### 📋 Historia de Usuario
```
Como equipo de operaciones
Queremos ver en qué ambiente está corriendo la API (development, staging, production)
Para poder identificar rápidamente el entorno al depurar problemas
```

#### ✅ Criterios de Aceptación

**CA1:** Agregar campo `environment` al endpoint `/status`
- Nuevo campo: `environment`
- Valor: `"development"` (por ahora hardcoded)
- Tipo: string

**CA2:** El campo debe aparecer en la respuesta JSON
```json
{
  "message": "...",
  "version": "1.1.0",
  "uptime": 5,
  "environment": "development"
}
```

#### 🧪 Escenarios de Prueba

**Escenario 1: Verificar nuevo campo**
```bash
curl http://localhost:8080/status
# Debe incluir: "environment": "development"
```

#### 💡 Pistas Técnicas

**Paso 1: Modificar el dominio**
- Archivo: `internal/domain/status/status.go`
- Agregar campo: `Environment string` con tag `` `json:"environment"` ``

**Paso 2: Asignar valor en el caso de uso**
- Archivo: `internal/usecase/status/get_status.go`
- Agregar: `Environment: "development",`

---

### EJERCICIO 7 - Agregar timestamp al status

#### 📋 Historia de Usuario
```
Como desarrollador
Quiero ver la fecha y hora actual del servidor en /status
Para verificar que el reloj del servidor esté sincronizado correctamente
```

#### ✅ Criterios de Aceptación

**CA1:** Agregar campo `timestamp` con fecha/hora actual
- Formato: ISO 8601 (RFC3339)
- Ejemplo: `"2025-12-15T10:30:45Z"`

**CA2:** El timestamp debe actualizarse en cada llamada
- Cada petición devuelve el momento exacto de la consulta
- No debe ser un valor fijo

**CA3:** Formato estándar internacional
- Usar `time.RFC3339` de Go
- Compatible con parsers de JavaScript, Python, etc.

#### 🧪 Escenarios de Prueba

**Escenario 1: Primera llamada**
```bash
curl http://localhost:8080/status
# Debe incluir: "timestamp": "2025-12-15T10:30:45Z" (hora actual)
```

**Escenario 2: Segunda llamada (unos segundos después)**
```bash
curl http://localhost:8080/status
# El timestamp debe ser diferente (más reciente)
```

#### 💡 Pistas Técnicas

**Paso 1: Agregar campo al dominio**
- Archivo: `internal/domain/status/status.go`
- Campo: `Timestamp string` con tag `json:"timestamp"`

**Paso 2: Generar timestamp en el caso de uso**
- Archivo: `internal/usecase/status/get_status.go`
- Usar: `time.Now().Format(time.RFC3339)`
- No olvides importar: `"time"`

---

### EJERCICIO 8 - Mensajes de error específicos

#### 📋 Historia de Usuario
```
Como usuario de la API
Quiero recibir mensajes de error claros y específicos
Para entender exactamente qué salió mal y cómo solucionarlo
```

#### ✅ Criterios de Aceptación

**CA1:** Diferenciar errores por código HTTP del servidor externo

| Código HTTP | Mensaje de Error |
|-------------|------------------|
| 404 | "usuario no encontrado" |
| 500-599 | "el servidor externo no está disponible" |
| Otros | "error inesperado del servidor externo" |

**CA2:** Mantener el código HTTP 400 para el cliente
- Aunque el externo devuelva 404, nosotros devolvemos 400
- Es un error de validación desde el punto de vista del cliente

**CA3:** Errores de conexión siguen manejándose como antes
- Timeout, DNS, etc. → Error de conexión genérico

#### 🧪 Escenarios de Prueba

**Escenario 1: Usuario no existe (404 del externo)**
```bash
curl http://localhost:8080/users/99
# Esperado: "usuario no encontrado"
```

**Escenario 2: Usuario válido**
```bash
curl http://localhost:8080/users/1
# Esperado: Datos del usuario (sin errores)
```

**Escenario 3: Simular servidor externo caído**
```bash
# (Este escenario es difícil de probar sin cambiar la URL)
# Pero tu código debe estar preparado para manejar 500+
```

#### 💡 Pistas Técnicas

**Archivo:** `internal/repository/user_api_repository.go`

**Opción 1: Usar switch**
```go
switch resp.StatusCode {
case http.StatusNotFound: // 404
    return nil, errors.New("usuario no encontrado")
case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
    return nil, errors.New("el servidor externo no está disponible")
default:
    return nil, fmt.Errorf("error inesperado del servidor externo: %d", resp.StatusCode)
}
```

**Opción 2: Usar if-else**
```go
if resp.StatusCode == 404 {
    return nil, errors.New("usuario no encontrado")
} else if resp.StatusCode >= 500 {
    return nil, errors.New("el servidor externo no está disponible")
} else {
    return nil, errors.New("error inesperado")
}
```

---

## 🟡 NIVEL INTERMEDIO - PARTE 1 (Ejercicios 9-14)

Crear endpoints completos, entender el flujo completo.

---

### EJERCICIO 9 - Listar todos los usuarios

#### 📋 Historia de Usuario
```
Como administrador de la tienda
Quiero ver la lista completa de usuarios registrados
Para poder gestionar la base de clientes y generar reportes
```

#### ✅ Criterios de Aceptación

**CA1:** Nuevo endpoint GET `/users` (sin ID)
- Devuelve array JSON con todos los usuarios
- No requiere parámetros
- HTTP 200 OK en caso exitoso

**CA2:** Formato de respuesta
```json
[
  {
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    ...
  },
  {
    "id": 2,
    "name": "Ervin Howell",
    ...
  }
]
```

**CA3:** El endpoint `/users/{id}` debe seguir funcionando
- No romper funcionalidad existente
- `/users/1` sigue devolviendo un solo usuario

**CA4:** Orden de las rutas importa
- `/users` debe registrarse ANTES de `/users/{id}` en el router
- Si no, `/users` podría interpretarse como un ID

#### 🧪 Escenarios de Prueba

**Escenario 1: Listar todos**
```bash
curl http://localhost:8080/users
# Esperado: Array con 10 usuarios
```

**Escenario 2: Verificar que usuario individual sigue funcionando**
```bash
curl http://localhost:8080/users/1
# Esperado: Un solo usuario (objeto, no array)
```

**Escenario 3: Verificar Content-Type**
```bash
curl -i http://localhost:8080/users
# Headers deben incluir: Content-Type: application/json
```

#### 💡 Pasos a seguir (Flujo completo):

1. **Actualizar interfaz del repositorio**
   - Archivo: `internal/domain/user/repository.go`
   - Agregar método: `FindAll() ([]*User, error)`

2. **Implementar en el repositorio**
   - Archivo: `internal/repository/user_api_repository.go`
   - Estudia `FindByID` primero
   - ¿Qué cambia para `FindAll`?
     - URL: `https://jsonplaceholder.typicode.com/users` (sin /%d)
     - Decode: `var users []*user.User` (slice, no puntero)
   
3. **Crear el caso de uso**
   - Archivo nuevo: `internal/usecase/user/list_users.go`
   - Estructura: `ListUsersUsecase`
   - Método: `Execute() ([]*user.User, error)`
   - Solo llama a `repository.FindAll()`
   
4. **Actualizar el handler**
   - Archivo: `internal/handler/user_handler.go`
   - Agregar campo: `listUsersUC *userUC.ListUsersUsecase`
   - Actualizar constructor para recibirlo
   - Crear método: `GetAll(w http.ResponseWriter, r *http.Request)`
     - Llamar al caso de uso
     - Devolver lista en JSON
   
5. **Conectar en main.go**
   - Crear instancia: `listUsersUC := userUC.NewListUsersUsecase(userRepo)`
   - Pasar al handler: `handler.NewUserHandler(getUserUC, listUsersUC)`
   
6. **Registrar ruta**
   - Archivo: `internal/router/router.go`
   - **IMPORTANTE:** Antes de `/users/{id}`
   - `r.Get("/users", userHandler.GetAll)`

**Prueba:** `curl http://localhost:8080/users`

---

### EJERCICIO 10 - Endpoint de bienvenida en la raíz

#### 📋 Historia de Usuario
```
Como nuevo desarrollador que descubre la API
Quiero ver información de bienvenida al acceder a la raíz (/)
Para conocer rápidamente qué endpoints están disponibles y la versión actual
```

#### ✅ Criterios de Aceptación

**CA1:** Nuevo endpoint GET `/` (raíz)
- Devuelve información de bienvenida
- HTTP 200 OK

**CA2:** Información a incluir
```json
{
  "message": "Bienvenido a la API de la Tienda Online",
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

**CA3:** Debe ser el primer endpoint que se encuentra
- Un desarrollador que hace `curl http://localhost:8080` ve esto
- Experiencia de "auto-documentación"

#### 🧪 Escenarios de Prueba

**Escenario 1: Acceder a la raíz**
```bash
curl http://localhost:8080/
# Esperado: JSON con bienvenida y lista de endpoints
```

**Escenario 2: Verificar que incluye versión**
```bash
curl http://localhost:8080/ | grep version
# Debe aparecer: "version": "1.1.0"
```

**Escenario 3: Desde navegador**
```
Abrir: http://localhost:8080/
# Debe mostrar JSON formateado
```

#### 💡 Pasos a seguir:

1. **Crear el dominio**
   - Carpeta: `internal/domain/welcome/`
   - Archivo: `welcome.go`
   - Estructura: `Welcome` con campos: `Message`, `Version`, `Endpoints []string`

2. **Crear el caso de uso**
   - Carpeta: `internal/usecase/welcome/`
   - Archivo: `get_welcome.go`
   - Devuelve: estructura Welcome con datos hardcoded

3. **Crear el handler**
   - Archivo: `internal/handler/welcome_handler.go`
   - Método: `Get(w http.ResponseWriter, r *http.Request)`

4. **Conectar todo**
   - En `main.go`: crear caso de uso y handler
   - En `router.go`: registrar `r.Get("/", welcomeHandler.Get)`

**Prueba:** `curl http://localhost:8080/`

---

### EJERCICIO 11 - Contador de peticiones al status

#### 📋 Historia de Usuario
```
Como administrador del sistema
Quiero ver cuántas veces se ha consultado el endpoint /status
Para monitorear la frecuencia de health checks de nuestros sistemas de monitoreo
```

#### ✅ Criterios de Aceptación

**CA1:** El endpoint `/status` debe incluir un campo `request_count` en la respuesta
- Tipo: número entero
- Valor inicial: 1 (primera llamada)

**CA2:** El contador debe incrementarse en 1 cada vez que se llama al endpoint
- Primera llamada: `request_count: 1`
- Segunda llamada: `request_count: 2`
- N llamada: `request_count: N`

**CA3:** El contador debe persistir mientras el servidor esté corriendo
- No debe reiniciarse entre llamadas
- SÍ debe reiniciarse al reiniciar el servidor (es aceptable por ahora)

#### 🧪 Escenarios de Prueba

**Escenario 1: Primera llamada**
```bash
curl http://localhost:8080/status
# Debe mostrar: "request_count": 1
```

**Escenario 2: Múltiples llamadas**
```bash
curl http://localhost:8080/status  # request_count: 1
curl http://localhost:8080/status  # request_count: 2
curl http://localhost:8080/status  # request_count: 3
```

**Escenario 3: Después de reiniciar servidor**
```bash
# 1. Hacer varias llamadas (contador llega a 5)
# 2. Ctrl+C para detener servidor
# 3. go run cmd/app/main.go (reiniciar)
# 4. curl http://localhost:8080/status
# Debe mostrar: request_count: 1 (reinició)
```

#### 💡 Pistas Técnicas

**Paso 1:** Modificar dominio (`internal/domain/status/status.go`)
- Agregar campo: `RequestCount int` con tag `json:"request_count"`

**Paso 2:** Modificar caso de uso (`internal/usecase/status/get_status.go`)
- Agregar campo privado a la estructura: `requestCount int`
- En el método `Execute()`: incrementar contador antes de crear respuesta
- Incluir contador en la respuesta

**Paso 3:** Verificar
- Reiniciar servidor
- Llamar al endpoint múltiples veces
- Verificar que el número incrementa

---

### EJERCICIO 12 - Obtener solo el nombre de un usuario

**Contexto de negocio:** Para el sistema de chat de la tienda, solo necesitamos mostrar el nombre del usuario, no toda su información completa.

**Problema:** Crear `GET /users/{id}/name` que devuelva solo `{"name": "Leanne Graham"}`

**Pistas:**
- Crear estructura simple en `internal/domain/user/user_name.go`:
  - Solo un campo: `Name string` con tag `json:"name"`
- En el handler: **reutilizar** el caso de uso `getUserUC` que ya existe
- Extraer solo el nombre del usuario completo
- NO necesitas crear un nuevo caso de uso

**Ejemplo de lo que debes hacer:**
```go
// En el handler, después de obtener el usuario:
user, err := h.getUserUC.Execute(id)
// Crear una estructura solo con el nombre:
response := UserName{Name: user.Name}
// Devolver en JSON
```

**Prueba:** `curl http://localhost:8080/users/1/name`

**Concepto clave:** Aprendes a crear "vistas parciales" de datos reutilizando lógica existente.

---

### EJERCICIO 13 - Obtener solo el email de un usuario

**Contexto de negocio:** Para el sistema de notificaciones, solo necesitamos el email del usuario.

**Problema:** Crear `GET /users/{id}/email` que devuelva solo `{"email": "Sincere@april.biz"}`

**Pistas:**
- Muy similar al ejercicio 12
- Crear estructura en `internal/domain/user/user_email.go`
- Campo: `Email string` con tag `json:"email"`
- Reutilizar `getUserUC` existente
- Extraer: `response := UserEmail{Email: user.Email}`

**Prueba:** `curl http://localhost:8080/users/1/email`

**Refuerzo:** Practicas el mismo patrón otra vez para internalizarlo.

---

### EJERCICIO 14 - Obtener solo el username de un usuario

**Contexto de negocio:** Para mostrar menciones en comentarios tipo "@username", necesitamos solo el username.

**Problema:** Crear `GET /users/{id}/username` que devuelva solo `{"username": "Bret"}`

**Pistas:**
- Mismo patrón que ejercicios 12 y 13
- Crear `internal/domain/user/user_username.go`
- Campo: `Username string` con tag `json:"username"`
- Reutilizar `getUserUC`
- Extraer: `Username: user.Username`

**Prueba:** `curl http://localhost:8080/users/1/username`

**Objetivo:** Reforzar el patrón de vistas parciales (lo has hecho 3 veces ya).

---

### EJERCICIO 15 - Buscar usuario por nombre exacto

**Contexto de negocio:** El administrador de la tienda necesita buscar clientes por su nombre completo para atención al cliente.

**Problema:** Crear `GET /users/search-by-name?name=Leanne Graham` que encuentre el usuario por nombre exacto.

**Diferencias con ejercicios anteriores:**
- Ahora usas **query parameter** (antes usabas path parameter)
- Necesitas **filtrar una lista** (antes solo devolvías un campo)

**Pistas muy detalladas:**

1. **En el handler:**
   - Extraer parámetro: `name := r.URL.Query().Get("name")`
   - ¿Está vacío? → Error 400
   - Llamar al nuevo caso de uso

2. **Crear caso de uso** `internal/usecase/user/search_by_name.go`:
   - Llamar a `repository.FindAll()` (reutilizar)
   - Recorrer la lista con un `for`
   - Comparar: `if user.Name == name`
   - Si encuentra → devolver ese usuario
   - Si termina el loop → error "usuario no encontrado"

3. **NO necesitas modificar el repositorio** (reutilizas `FindAll`)

**Prueba:** `curl "http://localhost:8080/users/search-by-name?name=Leanne Graham"`

**Concepto nuevo:** Query parameters + filtrado en memoria.

---

### EJERCICIO 16 - Buscar usuario por nombre parcial

**Contexto de negocio:** A veces el cliente de soporte no recuerda el nombre completo, solo parte de él.

**Problema:** Crear `GET /users/search-by-name-partial?name=Leanne` que encuentre usuarios cuyo nombre **contenga** ese texto.

**Pistas:**
- Muy similar al ejercicio 15
- La ÚNICA diferencia: en lugar de `==` usas `strings.Contains()`
- Importar: `"strings"`
- Comparar: `strings.Contains(user.Name, name)`
- Buscar case-insensitive: `strings.Contains(strings.ToLower(user.Name), strings.ToLower(name))`

**Prueba:** 
```bash
curl "http://localhost:8080/users/search-by-name-partial?name=leanne"  # Funciona con minúsculas
curl "http://localhost:8080/users/search-by-name-partial?name=Graham"  # Encuentra por apellido
```

**Refuerzo:** Practicas filtrado con una pequeña variación.

---

### EJERCICIO 17 - Buscar usuario por username

**Contexto de negocio:** Los usuarios se identifican por username único en la tienda. Necesitamos buscarlos así.

**Problema:** Crear `GET /users/search-by-username?username=Bret` que encuentre el usuario por username.

**Pistas:**
- EXACTAMENTE igual al ejercicio 15
- Solo cambias el campo: en lugar de `user.Name` usas `user.Username`
- Es repetitivo **a propósito** para que internalices el patrón

**Prueba:** `curl "http://localhost:8080/users/search-by-username?username=Bret"`

**Objetivo:** Reforzar el patrón de búsqueda (4ta vez que lo haces).

---

### EJERCICIO 18 - Buscar usuario por email

**Contexto de negocio:** Soporte técnico recibe emails de clientes y necesita encontrar su cuenta rápidamente.

**Problema:** Crear `GET /users/search-by-email?email=Sincere@april.biz`

**Pistas:**
- Mismo patrón que ejercicios 15-17
- Campo: `user.Email`
- Case-insensitive: `strings.EqualFold(user.Email, email)`

**Prueba:** `curl "http://localhost:8080/users/search-by-email?email=sincere@april.biz"`

**Objetivo:** 5ta repetición del mismo patrón (ya lo dominas).

---

### EJERCICIO 19 - Buscar usuarios por dominio de email

**Contexto de negocio:** Marketing quiere saber cuántos usuarios corporativos tenemos (emails de empresas, no gmail/hotmail).

**Problema:** Crear `GET /users/by-domain?domain=biz` que devuelva TODOS los usuarios con emails que terminen en ese dominio.

**Diferencia clave:** Ahora devuelves UNA LISTA, no un solo usuario.

**Pistas muy detalladas:**

1. **En el caso de uso:**
   - En lugar de devolver `*user.User` devuelves `[]*user.User`
   - En lugar de `return user, nil` cuando encuentras uno, lo agregas a una lista:
   ```go
   matchedUsers := []*user.User{}
   for _, user := range users {
       if strings.HasSuffix(user.Email, "@"+domain) {
           matchedUsers = append(matchedUsers, user)
       }
   }
   return matchedUsers, nil
   ```

2. **Lista vacía NO es error** - es válido que no haya usuarios con ese dominio

**Prueba:** 
```bash
curl "http://localhost:8080/users/by-domain?domain=biz"
# Debe devolver array con varios usuarios
```

**Concepto nuevo:** Devolver múltiples resultados de una búsqueda.

---

---

## 🟡 NIVEL INTERMEDIO - PARTE 2 (Ejercicios 20-27)

**Nueva funcionalidad:** Ahora la tienda venderá productos. Repetirás EXACTAMENTE los mismos patrones que aprendiste con usuarios, pero aplicados a productos.

---

### EJERCICIO 20 - Crear el dominio de productos

**Contexto de negocio:** La tienda necesita gestionar su catálogo de productos para venderlos online.

**Problema:** Crear las estructuras básicas para manejar productos (NO crear endpoints todavía).

**Pasos súper detallados:**

1. **Crear la entidad Product** en `internal/domain/product/product.go`:
   ```go
   package product

   type Product struct {
       ID          int     `json:"id"`
       Title       string  `json:"title"`
       Price       float64 `json:"price"`
       Description string  `json:"description"`
       Category    string  `json:"category"`
       Image       string  `json:"image"`
   }
   ```

2. **Crear la interfaz del repositorio** en `internal/domain/product/repository.go`:
   ```go
   package product

   type Repository interface {
       FindByID(id int) (*Product, error)
       FindAll() ([]*Product, error)
   }
   ```

3. **NO crear nada más por ahora** - esto es solo la base

**Verificación:** Los archivos se crean sin errores de compilación.

**Concepto:** Entiendes que primero defines QUÉ ES un producto (dominio) antes de implementar CÓMO obtenerlo.

---

### EJERCICIO 21 - Implementar repositorio de productos

**Contexto de negocio:** Los productos vienen de una API externa (`https://fakestoreapi.com`).

**Problema:** Implementar el repositorio que obtiene productos de la API externa.

**Pasos:**

1. **Crear** `internal/repository/product_api_repository.go`

2. **Copiar la estructura de** `user_api_repository.go` - es IDÉNTICO:
   - Misma estructura con `baseURL`
   - Mismo constructor
   - Mismos métodos `FindByID` y `FindAll`
   - Solo cambian: URL (`/products` en lugar de `/users`) y tipo (`Product` en lugar de `User`)

3. **URL de la API externa:** `https://fakestoreapi.com/products`

**Pistas:**
- `FindByID`: GET a `https://fakestoreapi.com/products/1`
- `FindAll`: GET a `https://fakestoreapi.com/products`
- Decodificar a `*product.Product` o `[]*product.Product`

**Verificación:** Compila sin errores (todavía no lo usamos).

**Objetivo:** Practicas crear un repositorio (2da vez).

---

### EJERCICIO 22 - Crear caso de uso para listar productos

**Contexto de negocio:** El catálogo de la tienda debe mostrar todos los productos disponibles.

**Problema:** Crear el caso de uso `ListProductsUsecase` (igual que `ListUsersUsecase`).

**Pasos:**

1. **Crear** `internal/usecase/product/list_products.go`
2. **Copiar estructura de** `internal/usecase/user/list_users.go`
3. Cambiar nombres: `User` → `Product`, `user` → `product`
4. Llamar a `repository.FindAll()` del repositorio de productos

**Pista:** Es literalmente copy-paste cambiando nombres.

**Verificación:** Compila sin errores.

---

### EJERCICIO 23 - Crear caso de uso para obtener producto por ID

**Contexto de negocio:** Los clientes hacen clic en un producto y necesitan ver sus detalles.

**Problema:** Crear `GetProductUsecase` (igual que `GetUserUsecase`).

**Pasos:**

1. **Crear** `internal/usecase/product/get_product.go`
2. **Copiar de** `internal/usecase/user/get_user.go`
3. Validar que ID esté entre 1 y 20 (la API tiene 20 productos)
4. Llamar a `repository.FindByID(id)`

**Pista:** Mismo patrón que con usuarios.

---

### EJERCICIO 24 - Crear handler de productos

**Contexto de negocio:** Necesitamos endpoints HTTP para que el frontend consuma los productos.

**Problema:** Crear `ProductHandler` con dos métodos: `List` y `GetByID`.

**Pasos:**

1. **Crear** `internal/handler/product_handler.go`
2. **Copiar estructura de** `user_handler.go`
3. Cambiar tipos: `UserHandler` → `ProductHandler`
4. Recibir los dos casos de uso (get y list)
5. Implementar:
   - `List(w, r)` - llama a `listProductsUC.Execute()`
   - `GetByID(w, r)` - extrae ID, valida, llama a `getProductUC.Execute(id)`

**Pista:** Es el mismo patrón que `UserHandler`, solo cambias los tipos.

---

### EJERCICIO 25 - Conectar productos en main.go

**Contexto de negocio:** Ahora conectamos todo para que funcione.

**Problema:** Agregar la creación de productos en `cmd/app/main.go`.

**Pasos muy detallados:**

1. **En la sección de repositorios:**
   ```go
   // Repositorios
   userRepo := repository.NewUserAPIRepository(cfg.ExternalAPIURL)
   productRepo := repository.NewProductAPIRepository("https://fakestoreapi.com")  // ← NUEVO
   ```

2. **En la sección de casos de uso:**
   ```go
   // Casos de uso de productos
   getProductUC := productUC.NewGetProductUsecase(productRepo)      // ← NUEVO
   listProductsUC := productUC.NewListProductsUsecase(productRepo)  // ← NUEVO
   ```

3. **En la sección de handlers:**
   ```go
   productHandler := handler.NewProductHandler(getProductUC, listProductsUC)  // ← NUEVO
   ```

4. **Pasar al router:**
   - Agregar `productHandler` como parámetro de `SetupRouter`

**Verificación:** Compila sin errores.

---

### EJERCICIO 26 - Registrar rutas de productos

**Contexto de negocio:** Exponer los productos al mundo.

**Problema:** Agregar las rutas en `internal/router/router.go`.

**Pasos:**

1. **Actualizar firma de SetupRouter** para recibir `productHandler`

2. **Registrar rutas:**
   ```go
   // Productos
   r.Get("/products", productHandler.List)
   r.Get("/products/{id}", productHandler.GetByID)
   ```

**Pruebas:**
```bash
curl http://localhost:8080/products          # Lista todos (20 productos)
curl http://localhost:8080/products/1        # Un producto
curl http://localhost:8080/products/999      # Error: ID fuera de rango
```

**¡CELEBRA!** ¡Acabas de crear un módulo completo! 🎉

**Concepto:** Ahora entiendes el flujo completo y puedes crear cualquier módulo nuevo.

---

### EJERCICIO 27 - Obtener solo el título de un producto

**Contexto de negocio:** Para un dropdown de selección, solo necesitamos ID y título, no toda la info del producto.

**Problema:** Crear `GET /products/{id}/title` que devuelva `{"id": 1, "title": "..."}`

**Pistas:**
- Crear `internal/domain/product/product_title.go`
- Estructura con: `ID int` y `Title string`
- Reutilizar `getProductUC` en el handler
- Extraer solo esos dos campos

**Prueba:** `curl http://localhost:8080/products/1/title`

**Objetivo:** Aplicas el patrón de vistas parciales que aprendiste con usuarios.

---

## 🟡 NIVEL INTERMEDIO - PARTE 3 (Ejercicios 28-33)

**Nueva habilidad:** Ahora aprenderás a filtrar y ordenar listas de datos.

---

### EJERCICIO 28 - Buscar productos por título exacto

**Contexto de negocio:** El cliente busca un producto específico escribiendo su nombre completo en el buscador.

**Problema:** Crear `GET /products/search?title=Fjallraven - Foldsack No. 1 Backpack` que encuentre el producto por título exacto.

**Pistas:**
- Mismo patrón que búsqueda de usuarios por nombre (ejercicio 15)
- Query parameter: `title`
- Llamar a `FindAll()` y filtrar con `==`
- Case-insensitive: `strings.EqualFold(p.Title, title)`

**Prueba:** `curl "http://localhost:8080/products/search?title=Fjallraven - Foldsack No. 1 Backpack"`

---

### EJERCICIO 29 - Buscar productos por título parcial

**Contexto de negocio:** El cliente no recuerda el nombre completo, solo escribe "shirt" en el buscador.

**Problema:** Crear `GET /products/search-partial?title=shirt` que encuentre productos cuyo título contenga esa palabra.

**Pistas:**
- Similar al ejercicio 16 (búsqueda parcial de usuarios)
- Usar: `strings.Contains(strings.ToLower(p.Title), strings.ToLower(title))`
- Devolver LISTA de productos (puede haber varios con "shirt")

**Prueba:** `curl "http://localhost:8080/products/search-partial?title=shirt"`

**Refuerzo:** 6ta vez que haces búsqueda por texto (ya lo dominas).

---

### EJERCICIO 30 - Filtrar productos por categoría

**Contexto de negocio:** La tienda tiene categorías (electronics, jewelery, men's clothing, women's clothing). Los clientes quieren ver solo productos de una categoría.

**Problema:** Crear `GET /products/category/{category}` que devuelva todos los productos de esa categoría.

**Pistas:**
- Path parameter: `chi.URLParam(r, "category")`
- Filtrar lista con `if p.Category == category`
- Case-insensitive: `strings.EqualFold(p.Category, category)`
- Devolver LISTA (hay múltiples productos por categoría)

**Prueba:** 
```bash
curl http://localhost:8080/products/category/electronics
curl http://localhost:8080/products/category/jewelery
```

---

### EJERCICIO 31 - Productos más caros que un precio

**Contexto de negocio:** Cliente dice "Muéstrame solo productos de más de $100".

**Problema:** Crear `GET /products/price-above?min=100` que devuelva productos con precio >= al especificado.

**Pistas:**
- Query parameter: `min`
- Convertir a número: `minPrice, err := strconv.ParseFloat(minStr, 64)`
- Validar que sea >= 0
- Filtrar: `if p.Price >= minPrice { ... }`

**Prueba:** `curl "http://localhost:8080/products/price-above?min=100"`

**Concepto nuevo:** Filtrado numérico (antes era solo texto).

---

### EJERCICIO 32 - Productos más baratos que un precio

**Contexto de negocio:** Cliente busca ofertas: "Solo productos de menos de $50".

**Problema:** Crear `GET /products/price-below?max=50`

**Pistas:**
- Exactamente igual al ejercicio 31
- Solo cambia la comparación: `if p.Price <= maxPrice { ... }`

**Prueba:** `curl "http://localhost:8080/products/price-below?max=50"`

**Refuerzo:** Practicas filtrado numérico (2da vez).

---

### EJERCICIO 33 - Productos en rango de precio

**Contexto de negocio:** Cliente quiere filtrar: "Entre $50 y $150".

**Problema:** Crear `GET /products/price-range?min=50&max=150`

**Pistas:**
- Dos query parameters: `min` y `max`
- Validar: `min <= max`
- Filtrar: `if p.Price >= min && p.Price <= max { ... }`

**Prueba:** `curl "http://localhost:8080/products/price-range?min=50&max=150"`

**Concepto:** Combinar dos condiciones (AND lógico).

---

## 🔴 NIVEL AVANZADO (Ejercicios 34-38)

**Nueva habilidad:** Ordenamiento, agregaciones y estadísticas.

---

### EJERCICIO 34 - Ordenar productos por precio (menor a mayor)

**Contexto de negocio:** Cliente quiere ver primero los productos más baratos.

**Problema:** Crear `GET /products/sorted-by-price` que ordene productos de menor a mayor precio.

**Pistas muy detalladas:**

1. **Importar:** `"sort"`

2. **En el handler** (después de obtener la lista):
   ```go
   products, err := h.listProductsUC.Execute()
   
   // Ordenar
   sort.Slice(products, func(i, j int) bool {
       return products[i].Price < products[j].Price
   })
   
   // Devolver lista ordenada
   ```

3. **Concepto:** El handler puede hacer transformaciones simples antes de devolver datos

**Prueba:** `curl http://localhost:8080/products/sorted-by-price`  
Verifica que el primero sea el más barato.

---

### EJERCICIO 35 - Ordenar productos por precio (mayor a menor)

**Contexto de negocio:** Cliente quiere ver primero los productos premium (más caros).

**Problema:** Crear `GET /products/sorted-by-price-desc`

**Pistas:**
- Exactamente igual al ejercicio 34
- Solo cambia: `<` por `>`
- `return products[i].Price > products[j].Price`

**Prueba:** `curl http://localhost:8080/products/sorted-by-price-desc`

**Refuerzo:** 2da vez que ordenas (ahora descendente).

---

### EJERCICIO 36 - Producto más barato

**Contexto de negocio:** Para la sección "Mejor Oferta" del sitio, mostramos el producto más barato.

**Problema:** Crear `GET /products/cheapest` que devuelva el producto con menor precio.

**Pistas:**
```go
products, err := h.listProductsUC.Execute()

cheapest := products[0]
for _, p := range products {
    if p.Price < cheapest.Price {
        cheapest = p
    }
}

// Devolver cheapest
```

**Prueba:** `curl http://localhost:8080/products/cheapest`

**Concepto:** Encontrar el mínimo en una lista.

---

### EJERCICIO 37 - Top 3 productos más caros

**Contexto de negocio:** Sección "Premium" muestra los 3 productos más exclusivos (caros).

**Problema:** Crear `GET /products/top-expensive` que devuelva los 3 productos más caros.

**Pistas:**
```go
products, err := h.listProductsUC.Execute()

// 1. Ordenar descendente (más caro primero)
sort.Slice(products, func(i, j int) bool {
    return products[i].Price > products[j].Price
})

// 2. Tomar solo los primeros 3
top3 := products[:3]

// 3. Devolver top3
```

**Prueba:** `curl http://localhost:8080/products/top-expensive`

**Concepto:** Ordenar + limitar resultados.

---

### EJERCICIO 38 - Estadísticas de la tienda

**Contexto de negocio:** Dashboard del administrador muestra métricas clave de la tienda.

**Problema:** Crear `GET /store/stats` que devuelva:
- Total de usuarios registrados
- Total de productos en catálogo
- Categorías disponibles
- Producto más barato
- Producto más caro

**Arquitectura - Handler orquesta:**

El handler llama a múltiples casos de uso y combina resultados:

```go
func (h *StatsHandler) Get(w http.ResponseWriter, r *http.Request) {
    // 1. Obtener usuarios
    users, _ := h.listUsersUC.Execute()
    
    // 2. Obtener productos
    products, _ := h.listProductsUC.Execute()
    
    // 3. Calcular estadísticas
    categories := extractUniqueCategories(products)
    cheapest := findCheapest(products)
    mostExpensive := findMostExpensive(products)
    
    // 4. Construir respuesta
    stats := struct {
        TotalUsers      int      `json:"total_users"`
        TotalProducts   int      `json:"total_products"`
        Categories      []string `json:"categories"`
        CheapestPrice   float64  `json:"cheapest_price"`
        MostExpensivePrice float64 `json:"most_expensive_price"`
    }{
        TotalUsers:         len(users),
        TotalProducts:      len(products),
        Categories:         categories,
        CheapestPrice:      cheapest.Price,
        MostExpensivePrice: mostExpensive.Price,
    }
    
    // 5. Devolver JSON
    json.NewEncoder(w).Encode(stats)
}
```

**Pasos:**
1. Crear `StatsHandler` que reciba ambos casos de uso (users y products)
2. Implementar funciones helper para calcular estadísticas
3. Conectar en main.go
4. Registrar ruta `/store/stats`

**Prueba:** `curl http://localhost:8080/store/stats`

**Concepto IMPORTANTE:** El handler puede orquestar múltiples casos de uso y hacer agregaciones simples. NO crear un UseCase para esto.

---

## ✅ PROYECTO COMPLETO

**Problema:** `GET /users/search?email=ejemplo@dominio.com`

**Diferencias clave:**
- **Query parameter** (no path parameter)
- Filtrar datos en memoria
- Reutilizar `FindAll()` (NO crear método nuevo en repositorio)

**Pasos:**

1. **Handler**
   - Extraer: `email := r.URL.Query().Get("email")`
   - Validar que no esté vacío (400)
   - Llamar al caso de uso

2. **Caso de uso**
   - Archivo: `internal/usecase/user/search_user_by_email.go`
   - Llamar a `repository.FindAll()`
   - Filtrar con loop:
     ```go
     for _, u := range users {
         if strings.EqualFold(u.Email, email) {
             return u, nil
         }
     }
     return nil, errors.New("usuario no encontrado")
     ```
   - Importar: `"strings"`

3. **Conectar**
   - Crear caso de uso en main.go
   - Pasar al handler
   - Ruta (ANTES de `/users/{id}`): `r.Get("/users/search", userHandler.SearchByEmail)`

**Prueba:** `curl "http://localhost:8080/users/search?email=Sincere@april.biz"`

---

### EJERCICIO 16 - Obtener solo el email de un usuario

**Problema:** `GET /users/{id}/email` que devuelva solo `{"email": "..."}`

**Pistas:**
- Crear estructura: `internal/domain/user/user_email.go`
- En el handler: reutilizar `getUserUC` existente
- Extraer solo el email y crear `UserEmail{Email: user.Email}`
- NO necesitas nuevo caso de uso

---

### EJERCICIO 17 - Filtrar productos por categoría

**Problema:** `GET /products/category/{category}` devuelve productos de esa categoría.

**Pistas:**
- Path parameter: `chi.URLParam(r, "category")`
- Caso de uso filtra con `FindAll()` + loop
- Comparar: `strings.EqualFold(p.Category, category)`

---

### EJERCICIO 18 - Productos con precio mayor a X

**Problema:** `GET /products/price-above?min=100`

**Pistas:**
- Query parameter: `r.URL.Query().Get("min")`
- Convertir a float64: `strconv.ParseFloat(minStr, 64)`
- Filtrar: `if product.Price >= minPrice { ... }`

---

### EJERCICIO 19 - Filtrar usuarios por dominio de email

**Problema:** `GET /users/by-domain?domain=gmail.com` devuelve usuarios con emails de ese dominio.

**Pistas:**
- Usar: `strings.HasSuffix(user.Email, "@"+domain)`
- Lista vacía `[]` es válida (no es error)

---

### EJERCICIO 20 - Listar solo títulos de productos

**Problema:** `GET /products/titles` devuelve `[{"id": 1, "title": "..."}, ...]`

**Pistas:**
- Crear: `internal/domain/product/product_title.go`
- En el handler: reutilizar `listProductsUC`
- Transformar en loop a `[]ProductTitle`

---

### EJERCICIO 21 - Contar productos por categoría

**Problema:** `GET /products/count-by-category` devuelve `[{"category": "electronics", "count": 6}, ...]`

**Pistas:**
- Crear mapa: `counts := make(map[string]int)`
- Recorrer productos: `counts[p.Category]++`
- Convertir mapa a slice de `CategoryCount`

---

### EJERCICIO 22 - Buscar productos por título

**Problema:** `GET /products/search?query=shirt` busca en el título.

**Pistas:**
- `strings.Contains(strings.ToLower(p.Title), strings.ToLower(query))`
- Case-insensitive

---

### EJERCICIO 23 - Ordenar productos por precio

**Problema:** `GET /products/by-price` ordena de menor a mayor precio.

**Pistas:**
- Importar: `"sort"`
- Usar:
  ```go
  sort.Slice(products, func(i, j int) bool {
      return products[i].Price < products[j].Price
  })
  ```

---

### EJERCICIO 24 - Ordenar productos por precio descendente

**Problema:** `GET /products/sorted-by-price` (como 23 pero del más caro al más barato).

**Pista:** Cambiar `<` por `>` en la función de sort.

---

### EJERCICIO 25 - Combinar filtros de categoría y precio

**Problema:** `GET /products/filter?category=electronics&minPrice=100`

**Pistas:**
- Obtener ambos parámetros
- Filtrar en dos pasos:
  - Primero por categoría (si no está vacío)
  - Luego por precio (si minPrice > 0)

---

### EJERCICIO 26 - Producto más barato

**Problema:** `GET /products/cheapest` devuelve el producto con precio menor.

**Pistas:**
- `cheapest := products[0]`
- Recorrer y actualizar si encuentras uno menor

---

### EJERCICIO 27 - Filtrar por rango de precio

**Problema:** `GET /products/in-range?min=50&max=150`

**Pistas:**
- Validar: `min <= max`
- Filtrar: `p.Price >= min && p.Price <= max`

---

### EJERCICIO 28 - Obtener ciudad de usuario

**Problema:** `GET /users/{id}/city` devuelve `{"city": "..."}`

**Pistas:**
- Reutilizar `getUserUC`
- Acceder a: `user.Address.City`

---

### EJERCICIO 29 - Rango de precio (alternativo)

**Problema:** Similar al 27, pero con nombre diferente: `GET /products/price-range?min=50&max=200`

---

### EJERCICIO 30 - Listar categorías únicas

**Problema:** `GET /products/categories` devuelve `{"categories": ["electronics", "jewelery", ...]}`

**Pistas:**
- Usar mapa como set: `categoryMap := make(map[string]bool)`
- `categoryMap[p.Category] = true`
- Convertir keys a slice

---

## 🔴 NIVEL AVANZADO (Ejercicios 31-38)

Orquestación, agregaciones, múltiples dominios.

---

### EJERCICIO 31 - Estadísticas generales

**Problema:** `GET /stats` muestra total de usuarios, productos, categorías y response time estimado.

**Arquitectura - Handler orquesta:**
- Llamar a `listUsersUC` → contar
- Llamar a `listProductsUC` → contar + extraer categorías únicas
- Construir respuesta agregada

**NO crear un UseCase para esto.** Es solo agregación simple.

---

### EJERCICIO 32 - Perfil completo de usuario

**Problema:** `GET /users/{id}/profile` devuelve usuario + total de usuarios + mensaje personalizado.

**Handler orquesta:**
- `getUserUC.Execute(id)`
- `listUsersUC.Execute()` → contar con `len()`
- Mensaje: `"Perfil de " + user.Name`

---

### EJERCICIO 33 - Health check detallado

**Problema:** `GET /health` verifica conectividad a APIs externas.

**Handler orquesta:**
- Intentar `getUserUC.Execute(1)` → verifica users API
- Intentar `getProductUC.Execute(1)` → verifica products API
- Estado:
  - Ambos OK → `"healthy"`
  - Solo uno → `"degraded"`
  - Ninguno → `"unhealthy"`

---

### EJERCICIO 34 - Resumen ejecutivo

**Problema:** `GET /summary` devuelve versión, total de recursos, estado, endpoints, timestamp.

**Handler orquesta:**
- `getStatusUC` → versión
- `listUsersUC` → contar
- `listProductsUC` → contar
- Total = users + products

---

### EJERCICIO 35 - Recomendaciones de productos

**Problema:** `GET /users/{id}/recommended-products` devuelve usuario + 3 productos aleatorios.

**Handler orquesta + selección aleatoria:**
```go
import "math/rand"

rand.Seed(time.Now().UnixNano())
for i := 0; i < 3; i++ {
    idx := rand.Intn(len(products))
    recommended = append(recommended, products[idx])
}
```

---

### EJERCICIO 36 - Comparación de productos

**Problema:** `GET /products/compare?ids=1,2,3` compara varios productos.

**Pistas:**
- `idsParam := r.URL.Query().Get("ids")`
- `strings.Split(idsParam, ",")`
- Llamar a `getProductUC` para cada ID
- Calcular: precio promedio, más barato, más caro

---

### EJERCICIO 37 - Dashboard principal

**Problema:** `GET /dashboard` con estado, totales y top 3 productos más caros.

**Handler orquesta:**
- Contar users y products
- Ordenar productos por precio descendente
- Tomar primeros 3: `products[:3]`

---

### EJERCICIO 38 - Búsqueda global

**Problema:** `GET /search?q=john` busca en usuarios (nombre, email, username) y productos (título, descripción).

**Handler orquesta:**
- Buscar en `listUsersUC` → filtrar
- Buscar en `listProductsUC` → filtrar
- Combinar resultados:
  ```go
  response := struct {
      MatchedUsers    []*user.User
      MatchedProducts []*product.Product
      TotalResults    int
  }{...}
  ```

---

¡FELICIDADES! Has construido una **API completa para una tienda online** con:

**📦 Funcionalidades implementadas:**
- ✅ Gestión de usuarios (registro, búsqueda, perfiles)
- ✅ Catálogo de productos (listado, búsqueda, filtros)
- ✅ Filtros avanzados (por precio, categoría, texto)
- ✅ Ordenamiento de productos
- ✅ Estadísticas y reportes
- ✅ Dashboard administrativo

**🏛️ Arquitectura profesional:**
- ✅ Clean Architecture implementada
- ✅ Separación en capas (Domain, UseCase, Handler, Repository)
- ✅ Inyección de dependencias
- ✅ Código escalable y mantenible

---

## ✅ CHECKLIST POR EJERCICIO

**Antes de continuar al siguiente:**

- [ ] ¿Funciona sin errores?
- [ ] ¿Probaste casos de error (IDs inválidos, parámetros vacíos)?
- [ ] ¿Entiendes CADA línea que escribiste?
- [ ] ¿Podrías explicárselo a alguien más sin mirar el código?
- [ ] ¿Podrías reescribirlo mañana de memoria?

**Si respondiste NO a alguna → NO sigas adelante.**  
Revisa, debuggea, entiende primero. Es mejor dominar 10 ejercicios que completar 38 sin entender.

---

## 📊 PROGRESIÓN DE APRENDIZAJE

### 🟢 Ejercicios 1-11: Fundamentos
- Modificar código existente
- Validaciones simples
- Agregar campos
- Entender el flujo básico

### 🟡 Ejercicios 12-19: Vistas Parciales y Búsquedas
- Reutilizar casos de uso
- Query parameters
- Búsqueda en listas
- Filtrado en memoria
- **Patrón repetido 8 veces** para internalizarlo

### 🟡 Ejercicios 20-27: Módulo Completo de Productos
- Crear dominio desde cero
- Implementar repositorio
- Crear casos de uso
- Conectar todo en main.go
- Registrar rutas
- **Aplicar patrones aprendidos a nuevo dominio**

### 🟡 Ejercicios 28-33: Filtros y Ordenamiento
- Filtrado por texto
- Filtrado por números
- Rangos de valores
- Ordenamiento (ascendente/descendente)
- Encontrar mínimos/máximos

### 🔴 Ejercicios 34-38: Funciones Avanzadas
- Operaciones con listas
- Top N elementos
- Orquestación de múltiples casos de uso
- Agregaciones y estadísticas
- Dashboard administrativo

---

## 🆘 SI TE ATORAS

### Paso 1: Diagnóstico (5 minutos)
1. **Lee el mensaje de error COMPLETO** (no solo la primera línea)
2. **Identifica dónde está el error** (archivo y línea)
3. **¿Qué intentabas hacer?** (describe en palabras simples)

### Paso 2: Comparación (10 minutos)
4. **Busca código similar** que sí funcione
5. **Compara línea por línea** - ¿qué es diferente?
6. **Revisa imports** - ¿falta alguno?
7. **Verifica nombres** (Go distingue mayúsculas/minúsculas)

### Paso 3: Google (10 minutos)
8. **Busca el error específico** en Google
9. **Busca tutoriales** del concepto que no entiendes
10. **Lee documentación oficial** de Go

### Paso 4: Reintentar (15 minutos)
11. **Reinicia el servidor** (Ctrl+C y vuelve a ejecutar)
12. **Comenta código** para aislar el problema
13. **Agrega `fmt.Println()`** para ver qué está pasando
14. **Prueba con curl** paso a paso

### Paso 5: Pedir Ayuda (si llevas 60+ minutos atascado)
- Prepara tu pregunta:
  - ¿Qué ejercicio estás haciendo?
  - ¿Qué intentabas lograr?
  - ¿Qué error obtienes? (copia completo)
  - ¿Qué ya intentaste?

**IMPORTANTE:** No pidas "dame el código". Pide ayuda para ENTENDER qué está mal.

---

## 🎯 CONCEPTOS CLAVE APRENDIDOS

### Flujo de Desarrollo (repite hasta que lo sueñes)
```
1. RUTA → Define QUÉ endpoint quieres
2. HANDLER → Maneja petición HTTP (extrae parámetros, valida formato)
3. CASO DE USO → Aplica lógica de negocio
4. REPOSITORIO → Obtiene/guarda datos
5. MAIN → Conecta todo con inyección de dependencias
```

### Convenciones de Nombres
```go
// ✅ BIEN - Claro y consistente
import userUC "ejercicio-api/internal/usecase/user"
import productUC "ejercicio-api/internal/usecase/product"

getUserUsecase := userUC.NewGetUserUsecase(repo)
listUsersUsecase := userUC.NewListUsersUsecase(repo)

// ❌ MAL - Confuso
import getUserUsecase "..."
usecase := getUserUsecase.NewUsecase()  // ¿Qué usecase?
```

### Validaciones (doble capa de defensa)
- **Handler:** Validaciones de formato
  - ¿Es un número? `strconv.Atoi()`
  - ¿Está vacío? `if param == ""`
  - ¿Formato válido? Regex, ParseFloat, etc.
  
- **UseCase:** Validaciones de negocio
  - ¿Está en rango? `if id < 1 || id > 10`
  - ¿Tiene permisos? (en apps reales)
  - ¿Cumple reglas de negocio?

### Reutilización vs Creación
**¿Cuándo REUTILIZAR casos de uso existentes?**
- ✅ Vistas parciales (ejercicios 12-14): solo extraes campos
- ✅ Transformaciones simples: formato, conteo, ordenamiento

**¿Cuándo CREAR nuevo caso de uso?**
- ✅ Nueva regla de negocio (búsquedas, filtros con lógica)
- ✅ Nueva validación de negocio
- ✅ Nuevo acceso al repositorio

### Handler Orquesta (ejercicios avanzados)
El Handler **SÍ puede:**
- ✅ Llamar a múltiples casos de uso
- ✅ Hacer agregaciones simples (contar, sumar)
- ✅ Transformar datos para presentación (ordenar, limitar)
- ✅ Combinar resultados de diferentes dominios

El Handler **NO debe:**
- ❌ Implementar lógica de negocio compleja
- ❌ Acceder directamente al repositorio
- ❌ Manejar transacciones de BD

Los UseCases **NO deben:**
- ❌ Llamar a otros UseCases (eso es orquestación = trabajo del Handler)
- ❌ Conocer detalles de HTTP
- ❌ Devolver JSON (eso es trabajo del Handler)

---

## 🎓 HABILIDADES DOMINADAS

Si completaste los 38 ejercicios **entendiendo cada uno**, ahora dominas:

### Técnicas
- ✅ Desarrollo de APIs REST completas
- ✅ Clean Architecture en práctica
- ✅ Inyección de dependencias
- ✅ Separación de responsabilidades
- ✅ Patrones de diseño aplicados
- ✅ Manejo de errores apropiado
- ✅ Validaciones en múltiples capas
- ✅ Query parameters y path parameters
- ✅ Filtrado y búsqueda en memoria
- ✅ Ordenamiento de colecciones
- ✅ Agregaciones y estadísticas

### Arquitectura
- ✅ Diseño por capas
- ✅ Domain-Driven Design (DDD)
- ✅ Hexagonal Architecture
- ✅ Repositorios como abstracciones
- ✅ Casos de uso independientes
- ✅ Handlers como adaptadores

### Go Específico
- ✅ Estructuras y tags JSON
- ✅ Interfaces y contratos
- ✅ Manejo de errores con múltiples valores de retorno
- ✅ Slices y maps
- ✅ Filtrado con loops
- ✅ Ordenamiento con `sort.Slice()`
- ✅ Manipulación de strings
- ✅ HTTP requests y JSON encoding/decoding

---

## 🚀 PRÓXIMOS PASOS

### Nivel 1: Reforzar (recomendado)
1. **Rehaz ejercicios clave** sin mirar tus notas:
   - Ejercicio 9: Listar usuarios
   - Ejercicio 14: Módulo de productos
   - Ejercicio 28: Búsqueda con filtros
   - Ejercicio 38: Estadísticas

2. **Crea variaciones:**
   - Agrega módulo de "Carrito de compras"
   - Implementa "Favoritos" por usuario
   - Crea sistema de "Reviews" de productos

### Nivel 2: Expandir
3. **Persistencia real:**
   - Cambia APIs externas por PostgreSQL
   - Implementa repositorio con SQL
   - Aprende sobre transacciones

4. **Autenticación:**
   - Implementa JWT tokens
   - Protege endpoints privados
   - Roles y permisos

5. **Testing:**
   - Unit tests de casos de uso
   - Integration tests de handlers
   - Mocks de repositorios

### Nivel 3: Profesionalizar
6. **Documentación:** Swagger/OpenAPI
7. **Observabilidad:** Logs, métricas, tracing
8. **Deployment:** Docker, Kubernetes
9. **CI/CD:** GitHub Actions, tests automáticos
10. **Performance:** Caché, rate limiting, paginación

---

## 💪 MENSAJE FINAL

**¡FELICIDADES por llegar hasta aquí!** 🎉

Has construido algo real y funcional. No solo "completaste ejercicios", creaste una **API para una tienda online** con arquitectura profesional.

**Recuerda:**
- ✅ La velocidad NO importa
- ✅ La comprensión SÍ importa
- ✅ Mejor 10 ejercicios dominados que 38 copiados
- ✅ El código que entiendes es el código que puedes mantener
- ✅ Los patrones que internalizas son tu superpoder

**Ahora puedes:**
- Crear APIs REST desde cero
- Diseñar arquitecturas limpias
- Trabajar en proyectos Go profesionales
- Explicar Clean Architecture a otros
- Escalar sistemas de manera sostenible

**Comparte tu logro:** Si completaste el taller, compártelo en redes. Ayuda a otros aprendices compartiendo tu experiencia.

**¡Sigue construyendo cosas increíbles!** 🚀

---

## 📝 CERTIFICADO DE FINALIZACIÓN

Si completaste los 38 ejercicios, **¡lo lograste!** Copia este certificado en tu README personal:

```
🎓 CERTIFICADO DE FINALIZACIÓN

He completado el taller de 38 ejercicios de APIs REST con Clean Architecture en Go.

✅ 38 ejercicios completados
✅ API de tienda online funcional
✅ Clean Architecture implementada
✅ Patrones de diseño aplicados

Fecha de finalización: [TU FECHA AQUÍ]
Proyecto: [LINK A TU REPO]

#golang #cleanarchitecture #api #rest
```

¡Adelante y que sigas aprendiendo! 💪


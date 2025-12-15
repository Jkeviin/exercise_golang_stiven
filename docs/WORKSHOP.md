# 🎓 TALLER PRÁCTICO - Desarrollo de APIs REST

Este taller te guiará para agregar nuevas funcionalidades a la API paso a paso.

**Importante**: Cada ejercicio explica QUÉ necesitas agregar desde la perspectiva de negocio.

---

## 🚀 ANTES DE EMPEZAR

### Verifica que todo funciona

1. **Inicia el servidor**
   - Abre una terminal
   - Ve a la carpeta del proyecto
   - Ejecuta: `go run cmd/app/main.go`
   - Deberías ver: "🚀 Servidor iniciado en http://localhost:8080"

2. **Prueba los endpoints existentes**
   - Abre otra terminal (deja el servidor corriendo)
   - Ejecuta: `curl http://localhost:8080/status`
   - Deberías ver información del servidor
   - Ejecuta: `curl http://localhost:8080/ping`
   - Deberías ver: `{"message":"pong"}`
   - Ejecuta: `curl http://localhost:8080/users/1`
   - Deberías ver información de un usuario

Si todo funciona, ¡estás listo para empezar! 🎉

---

## 📚 ENTENDIENDO EL FLUJO DE DESARROLLO

Antes de empezar con los ejercicios, es importante entender el **orden cronológico** en el que creamos un endpoint:

### 🔄 FLUJO CORRECTO: Ruta → Handler → Caso de Uso → Repositorio

```
1. RUTA (router.go)
   ↓ "Quiero un endpoint /users"
   
2. HANDLER (handler/)
   ↓ "Recibe la petición HTTP, extrae parámetros"
   
3. CASO DE USO (usecase/)
   ↓ "Aplica lógica de negocio y validaciones"
   
4. REPOSITORIO (adapter/repository/)
   ↓ "Obtiene los datos de APIs externas o BD"
   
5. DOMINIO (domain/)
   "Define las estructuras de datos y contratos"
```

### ✅ ¿Por qué este orden?

1. **Empezamos por la Ruta** porque primero definimos QUÉ queremos exponer
2. **Creamos el Handler** que recibirá las peticiones HTTP
3. **Creamos el Caso de Uso** con la lógica de negocio
4. **Definimos el Repositorio** si necesitamos obtener datos
5. **El Dominio** define las estructuras y contratos que todos usan

### 📝 Tipos de Ejercicios

**🟢 Ejercicios Básicos (1-8)**: Modificar código existente
- Solo cambias valores o agregas validaciones
- No creas archivos nuevos

**🟡 Ejercicios Intermedios (9-30)**: Crear endpoints simples
- Sigues el flujo: Ruta → Handler → Caso de Uso → Repositorio
- Puedes reutilizar casos de uso existentes
- Transformaciones en memoria (filtros, ordenamiento)

**🔴 Ejercicios Avanzados (31-38)**: Endpoints complejos
- Combinas múltiples casos de uso
- Integras diferentes dominios
- Lógica de negocio más compleja

### 💡 Consejos para cada ejercicio

1. **Lee TODO el ejercicio** antes de empezar a escribir código
2. **Sigue los pasos en orden** - están numerados por una razón
3. **No te saltes pasos** - cada uno construye sobre el anterior
4. **Prueba frecuentemente** - reinicia el servidor después de cada cambio
5. **Si algo falla**, revisa los pasos anteriores antes de continuar

### 🎯 Ejemplo Visual del Flujo

```go
// 1. RUTA (comentada primero)
// r.Get("/users", userHandler.List)

// 2. HANDLER
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
    users, err := h.listUsersUsecase.Execute()  // Llama al caso de uso
    // ... manejo de respuesta
}

// 3. CASO DE USO
func (uc *ListUsersUsecase) Execute() ([]*User, error) {
    return uc.repository.FindAll()  // Llama al repositorio
}

// 4. REPOSITORIO
func (r *UserAPIRepository) FindAll() ([]*User, error) {
    // Hace petición HTTP a API externa
    // Devuelve los datos
}

// 5. DOMINIO (ya definido)
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### ⚠️ IMPORTANTE: Cuándo usar Handler vs UseCase

**📌 REGLA DE ORO:**

**Handler orquesta → UseCase ejecuta lógica de negocio → Repository obtiene datos**

#### ✅ El HANDLER debe:
- ✅ Orquestar llamadas a **múltiples** UseCases
- ✅ Combinar resultados de diferentes fuentes
- ✅ Hacer transformaciones simples (formateo, conteo, ordenamiento)
- ✅ Aplicar filtros en memoria
- ✅ Construir respuestas HTTP específicas
- ✅ Manejar parámetros de query/URL

#### ✅ El USECASE debe:
- ✅ Implementar **UNA** regla de negocio específica
- ✅ Validar datos de negocio
- ✅ Llamar al Repositorio para obtener/guardar datos
- ✅ Aplicar lógica de dominio compleja
- ✅ Ser **independiente** de otros UseCases
- ✅ Ser **reutilizable** desde diferentes Handlers

#### ❌ El USECASE NO debe:
- ❌ Llamar a otros UseCases (eso es orquestación = responsabilidad del Handler)
- ❌ Conocer detalles de HTTP (códigos de estado, headers, JSON)
- ❌ Combinar datos solo para presentación (el Handler lo hace)

#### 🤔 ¿Cuándo SÍ crear un UseCase que use otro UseCase?

**SOLO** si representa una **transacción atómica de negocio**:

```go
// ✅ VÁLIDO: Transferencia bancaria (transacción atómica)
type TransferMoneyUsecase struct {
    debitUC  *DebitAccountUsecase
    creditUC *CreditAccountUsecase
}
// Razón: "Transferir dinero" ES una regla de negocio que debe ser atómica
```

```go
// ❌ INVÁLIDO: Combinar usuario + status (solo presentación)
type GetUserInfoUsecase struct {
    getUserUC   *GetUserUsecase
    getStatusUC *GetStatusUsecase
}
// Razón: No es una regla de negocio, es solo combinar datos para el frontend
// Solución: El Handler lo hace directamente
```

### ⚠️ IMPORTANTE: Antes de empezar

- ✅ Asegúrate de entender el flujo Ruta → Handler → Caso de Uso → Repositorio
- ✅ Entiende **cuándo el Handler orquesta** vs **cuándo crear un UseCase**
- ✅ Lee cada ejercicio COMPLETO antes de escribir código
- ✅ Sigue los pasos numerados en orden
- ✅ Prueba después de cada paso importante
- ✅ Si te atoras, revisa los pasos previos y el ejemplo del flujo arriba

### 🆘 ¿Qué hacer si algo no funciona?

1. **Revisa los mensajes de error** - Go te dice exactamente qué falta
2. **Verifica imports** - ¿Importaste los paquetes necesarios?
3. **Revisa nombres** - Go distingue mayúsculas/minúsculas
4. **Reinicia el servidor** - Siempre necesario después de cambios
5. **Compara con código existente** - Los ejercicios iniciales siguen patrones similares

---

## EJERCICIO 1 - Cambiar el mensaje de bienvenida del ping

### 📋 LO QUE NECESITAMOS

El cliente dice que cuando alguien llame al endpoint `/ping`, el mensaje actual `"pong"` no es claro. Necesitamos que diga `"API funcionando correctamente"` para que sea más descriptivo.

### 🎯 INSTRUCCIONES

1. **Ubica donde está definido el mensaje del ping**
   - Ve a la carpeta: `internal/usecase/ping/`
   - Abre el archivo que contiene el texto `"pong"`

2. **Cambia el texto**
   - Donde dice `"pong"`, cámbialo por `"API funcionando correctamente"`

3. **Verifica el cambio**
   - Detén el servidor (Ctrl+C en la terminal)
   - Vuelve a iniciarlo: `go run cmd/app/main.go`
   - En otra terminal prueba: `curl http://localhost:8080/ping`
   - Ahora debe mostrar: `{"message":"API funcionando correctamente"}`

### ✅ RESULTADO ESPERADO

Antes: `{"message":"pong"}`
Ahora: `{"message":"API funcionando correctamente"}`

### 💡 LO QUE HICISTE

Modificaste el mensaje de respuesta de un endpoint para que sea más claro para los usuarios.

---

## EJERCICIO 2 - Cambiar el número de versión

### 📋 LO QUE NECESITAMOS

Acabamos de hacer una mejora en la aplicación. El equipo de producto solicita que actualicemos el número de versión que muestra el endpoint `/status` de `"1.0.0"` a `"1.1.0"`.

### 🎯 INSTRUCCIONES

1. **Ubica donde está definida la versión**
   - Ve a la carpeta: `internal/usecase/status/`
   - Busca donde dice `Version: "1.0.0"`

2. **Actualiza el número**
   - Cambia `"1.0.0"` por `"1.1.0"`

3. **Verifica el cambio**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - El campo `version` debe mostrar `"1.1.0"`

### ✅ RESULTADO ESPERADO

El endpoint `/status` ahora muestra la versión `1.1.0`.

### 💡 LO QUE HICISTE

Actualizaste un valor que devuelve el endpoint para reflejar cambios en la aplicación.

---

## EJERCICIO 3 - Rechazar IDs de usuario negativos

### 📋 LO QUE NECESITAMOS

Un usuario reportó que puede llamar `/users/-5` y el servidor intenta buscar ese usuario. Necesitamos que cuando alguien envíe un ID negativo o cero, el servidor responda inmediatamente con un error que diga: "El ID debe ser un número positivo".

### 🎯 INSTRUCCIONES

1. **Ubica donde se recibe el ID del usuario**
   - Ve a la carpeta: `internal/usecase/user/`
   - Abre el archivo donde se procesa el ID
   - Busca donde dice `if id <= 0`

2. **Asegúrate que existe la validación**
   - Debe haber una línea que verifica: `if id <= 0`
   - Si existe, el trabajo ya está hecho
   - Si no existe, agrégala con el mensaje: `"el ID debe ser mayor que 0"`

3. **Prueba los casos**
   - ID positivo: `curl http://localhost:8080/users/5` → Debe funcionar
   - ID cero: `curl http://localhost:8080/users/0` → Debe dar error
   - ID negativo: `curl http://localhost:8080/users/-1` → Debe dar error

### ✅ RESULTADO ESPERADO

- `/users/5` → ✅ Funciona (muestra el usuario)
- `/users/0` → ❌ Error: "el ID debe ser mayor que 0"
- `/users/-1` → ❌ Error: "el ID debe ser mayor que 0"

### 💡 LO QUE HICISTE

Agregaste una protección para que no se procesen IDs inválidos.

---

## EJERCICIO 4 - Limitar los IDs de usuario a un rango válido

### 📋 LO QUE NECESITAMOS

El API externa que usamos solo tiene usuarios del 1 al 10. Sin embargo, alguien puede llamar `/users/999` y el servidor intenta buscarlo, lo cual tarda tiempo y falla. 

Necesitamos que si alguien pide un usuario con ID mayor a 10, el servidor responda inmediatamente con: "El ID debe estar entre 1 y 10".

### 🎯 INSTRUCCIONES

1. **Ubica la validación de IDs**
   - En `internal/usecase/user/`
   - Encuentra donde validas `if id <= 0`

2. **Agrega una segunda validación**
   - Después de validar que sea mayor que 0
   - Agrega otra validación que verifique si el ID es mayor que 10
   - El error debe decir: `"el ID debe estar entre 1 y 10"`

3. **Prueba todos los casos**
   - ID 5: `curl http://localhost:8080/users/5` → Debe funcionar
   - ID 10: `curl http://localhost:8080/users/10` → Debe funcionar
   - ID 11: `curl http://localhost:8080/users/11` → Debe dar error
   - ID 999: `curl http://localhost:8080/users/999` → Debe dar error

### ✅ RESULTADO ESPERADO

- `/users/1` a `/users/10` → ✅ Funcionan
- `/users/11` o mayor → ❌ Error: "el ID debe estar entre 1 y 10"
- `/users/0` o negativo → ❌ Error: "el ID debe ser mayor que 0"

### 💡 LO QUE HICISTE

Definiste un rango válido de valores para evitar consultas innecesarias.

---

## EJERCICIO 5 - Rechazar IDs que no sean números

### 📋 LO QUE NECESITAMOS

Algunos usuarios intentan llamar `/users/abc` o `/users/hola`. El sistema debe responder inmediatamente con un error claro: "El ID debe ser un número válido".

### 🎯 INSTRUCCIONES

1. **Ubica donde se procesa el parámetro de la URL**
   - Ve a la carpeta: `internal/adapter/http/handler/`
   - Abre el archivo `user_handler.go`
   - Busca donde se convierte el ID a número (función `Atoi`)

2. **Verifica que se maneja el error**
   - Ya debe haber código que maneja si la conversión falla
   - Busca `if err != nil` después de la conversión
   - El mensaje debe decir algo como "ID inválido"

3. **Prueba los casos**
   - ID válido: `curl http://localhost:8080/users/5` → Debe funcionar
   - ID texto: `curl http://localhost:8080/users/abc` → Debe dar error
   - ID especial: `curl http://localhost:8080/users/@#$` → Debe dar error

### ✅ RESULTADO ESPERADO

- `/users/5` → ✅ Funciona
- `/users/abc` → ❌ Error: "ID inválido"
- `/users/xyz123` → ❌ Error: "ID inválido"

### 💡 LO QUE HICISTE

Te aseguraste de que solo se procesen números válidos como IDs.

---

## EJERCICIO 6 - Agregar un campo de ambiente al status

### 📋 LO QUE NECESITAMOS

El equipo de operaciones necesita saber en qué ambiente está corriendo la aplicación. Queremos que el endpoint `/status` incluya un nuevo campo llamado `environment` que diga `"development"`.

### 🎯 INSTRUCCIONES

1. **Agrega el campo a la estructura**
   - Ve a: `internal/domain/status/`
   - Abre el archivo de Status
   - Verás campos como `Message`, `Version`, `Uptime`
   - Agrega un nuevo campo: `Environment` de tipo `string` con tag JSON `"environment"`

2. **Haz que se devuelva el valor**
   - Ve a: `internal/usecase/status/`
   - Busca donde se crean los valores de Status
   - Agrega: `Environment: "development"`

3. **Verifica el cambio**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - Debe aparecer: `"environment":"development"`

### ✅ RESULTADO ESPERADO

La respuesta incluye:
```json
{
  "message": "...",
  "version": "1.1.0",
  "uptime": 5,
  "environment": "development"
}
```

### 💡 LO QUE HICISTE

Agregaste información adicional que el equipo necesita ver.

---

## EJERCICIO 7 - Agregar fecha y hora al status

### 📋 LO QUE NECESITAMOS

Para debugging, necesitamos que el endpoint `/status` también muestre la fecha y hora actual del servidor. Agrégala en un campo llamado `timestamp`.

### 🎯 INSTRUCCIONES

1. **Agrega el campo timestamp**
   - En `internal/domain/status/`
   - Agrega campo: `Timestamp` de tipo `string` con tag JSON `"timestamp"`

2. **Genera la fecha/hora actual**
   - En `internal/usecase/status/`
   - Donde creas el Status, agrega: `Timestamp: time.Now().Format(time.RFC3339)`
   - Si marca error, agrega `"time"` en los imports del archivo

3. **Verifica**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - Debe aparecer algo como: `"timestamp":"2025-12-11T14:30:45Z"`

### ✅ RESULTADO ESPERADO

Ahora incluye la fecha/hora:
```json
{
  "message": "...",
  "version": "1.1.0",
  "uptime": 3,
  "environment": "development",
  "timestamp": "2025-12-11T14:30:45Z"
}
```

### 💡 LO QUE HICISTE

Agregaste información temporal útil para monitoreo.

---

## EJERCICIO 8 - Mensaje específico cuando un usuario no existe

### 📋 LO QUE NECESITAMOS

Cuando alguien pide un usuario que no existe (ejemplo: `/users/99`), el error actual no es claro. Necesitamos que cuando el servidor externo responda con error 404, nuestro sistema devuelva: "Usuario no encontrado".

### 🎯 INSTRUCCIONES

1. **Ubica donde se llama al servidor externo**
   - Ve a: `internal/adapter/repository/`
   - Abre `user_api_repository.go`
   - Busca donde se valida `resp.StatusCode`

2. **Agrega manejo específico para 404**
   - Reemplaza la validación actual por una que detecte diferentes códigos:
   - Si el código es 404: devolver error "usuario no encontrado"
   - Si el código es 500 o mayor: devolver error "el servidor externo no está disponible"
   - Para otros códigos: devolver "error inesperado del servidor"

3. **Prueba**
   - Usuario válido: `curl http://localhost:8080/users/1` → Funciona
   - Usuario inexistente: `curl http://localhost:8080/users/99` → Error claro

### ✅ RESULTADO ESPERADO

- `/users/1` → ✅ Muestra el usuario
- `/users/99` → ❌ "usuario no encontrado"

### 💡 LO QUE HICISTE

Mejoraste los mensajes de error para que sean más claros para los usuarios.

---

## EJERCICIO 9 - Crear endpoint para listar todos los usuarios

### 📋 LO QUE NECESITAMOS

Actualmente solo podemos ver un usuario a la vez con `/users/1`, `/users/2`, etc. El cliente necesita un nuevo endpoint `/users` (sin ID) que devuelva la lista completa de usuarios disponibles.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta (Empezamos desde aquí)**

1. Ve a: `internal/infrastructure/http/router.go`
2. Busca donde está la ruta `/users/{id}`
3. Arriba de esa línea, **comenta temporalmente** la nueva ruta que crearás:
   ```go
   // r.Get("/users", userHandler.List)
   ```
4. Esto te ayudará a recordar qué estás construyendo

**PASO 2: Crear el método en el Handler**

5. Ve a: `internal/adapter/http/handler/user_handler.go`
6. Agrega un nuevo campo en la estructura del handler para el caso de uso:
   ```go
   listUsersUsecase *user.ListUsersUsecase
   ```
7. Actualiza el constructor para recibir este caso de uso
8. Crea el método `List`:
   - Llama al caso de uso
   - Si hay error, devuelve 500
   - Si todo bien, devuelve la lista en JSON

**PASO 3: Crear el Caso de Uso**

9. Ve a: `internal/usecase/user/`
10. Crea archivo nuevo: `list_users.go`
11. Crea la estructura `ListUsersUsecase` que necesita el repositorio
12. Crea el método `Execute()` que:
    - Llama al repositorio para obtener todos los usuarios
    - Devuelve la lista o el error

**PASO 4: Actualizar el contrato del Repositorio**

13. Ve a: `internal/domain/user/repository.go`
14. Agrega el método: `FindAll() ([]*User, error)`

**PASO 5: Implementar en el Repositorio**

15. Ve a: `internal/adapter/repository/user_api_repository.go`
16. Implementa el método `FindAll`:
    - Construye la URL: `{baseURL}/users` (sin ID)
    - Haz la petición HTTP GET
    - Decodifica la respuesta en un slice de usuarios
    - Maneja los errores apropiadamente

**PASO 6: Conectar todo en el Router**

17. Vuelve a: `internal/infrastructure/http/router.go`
18. Crea la instancia del caso de uso:
    ```go
    listUsersUsecase := user.NewListUsersUsecase(userRepo)
    ```
19. Pásalo al handler al crearlo
20. Descomenta y activa la ruta: `r.Get("/users", userHandler.List)`

**PASO 7: Prueba**

21. Reinicia el servidor
22. Ejecuta: `curl http://localhost:8080/users`
23. Debes ver una lista de 10 usuarios

### ✅ RESULTADO ESPERADO

- `/users` → Lista completa (array de usuarios)
- `/users/1` → Sigue funcionando (un solo usuario)

### 💡 LO QUE HICISTE

Creaste un endpoint completo nuevo siguiendo el flujo natural: Ruta → Handler → Caso de Uso → Repositorio.

---

## EJERCICIO 10 - Endpoint de bienvenida en la raíz

### 📋 LO QUE NECESITAMOS

Cuando alguien entra a `http://localhost:8080/` queremos mostrar un mensaje de bienvenida con información básica:
- Un mensaje amigable
- La versión de la API
- Lista de endpoints disponibles

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. Ve a: `internal/infrastructure/http/router.go`
2. Al inicio de las rutas, comenta temporalmente:
   ```go
   // r.Get("/", welcomeHandler.Get)
   ```

**PASO 2: Crear el Handler**

3. En `internal/adapter/http/handler/`
4. Crea archivo: `welcome_handler.go`
5. Crea la estructura `WelcomeHandler` con el caso de uso
6. Crea el método `Get` que:
   - Llama al caso de uso
   - Devuelve el resultado en JSON

**PASO 3: Crear el Caso de Uso**

7. Crea carpeta: `internal/usecase/welcome/`
8. Crea archivo: `get_welcome.go`
9. Crea `GetWelcomeUsecase` que en su método `Execute()` devuelva:
   - Message: "Bienvenido a la API de Ejercicio"
   - Version: "1.1.0"
   - Endpoints: ["/status", "/ping", "/users", "/users/{id}"]

**PASO 4: Crear el Dominio**

10. Crea carpeta: `internal/domain/welcome/`
11. Crea archivo: `welcome.go`
12. Define la estructura `Welcome` con: `Message`, `Version`, `Endpoints` (slice de strings)

**PASO 5: Conectar en el Router**

13. Vuelve a `router.go`
14. Crea la instancia del caso de uso
15. Crea el handler pasándole el caso de uso
16. Activa la ruta: `r.Get("/", welcomeHandler.Get)`

**PASO 6: Prueba**

17. Reinicia el servidor
18. Ejecuta: `curl http://localhost:8080/`
19. Debe mostrar la bienvenida

### ✅ RESULTADO ESPERADO

```json
{
  "message": "Bienvenido a la API de Ejercicio",
  "version": "1.1.0",
  "endpoints": ["/status", "/ping", "/users", "/users/{id}"]
}
```

### 💡 LO QUE HICISTE

Creaste un endpoint completo desde la ruta hasta el dominio, siguiendo el flujo natural de desarrollo.

---

## EJERCICIO 11 - Contador de peticiones al status

### 📋 LO QUE NECESITAMOS

Para monitoreo, necesitamos saber cuántas veces se ha llamado al endpoint `/status`. Agrégale un campo `request_count` que se incremente en cada llamada.

### 🎯 INSTRUCCIONES

1. **Agrega el campo**
   - En `internal/domain/status/`
   - Agrega: `RequestCount` de tipo `int` con tag JSON `"request_count"`

2. **Implementa el contador**
   - En `internal/usecase/status/`
   - Agrega un campo `requestCount int` a la estructura del caso de uso
   - En el método Execute, incrementa: `uc.requestCount++`
   - Incluye en la respuesta: `RequestCount: uc.requestCount`

3. **Prueba**
   - Llama varias veces: `curl http://localhost:8080/status`
   - El número debe ir aumentando: 1, 2, 3, etc.

### ✅ RESULTADO ESPERADO

Primera llamada: `"request_count": 1`
Segunda llamada: `"request_count": 2`
Tercera llamada: `"request_count": 3`

### 💡 LO QUE HICISTE

Agregaste una métrica simple de uso del endpoint.

---

## EJERCICIO 12 - Endpoint que combina usuario y estado del servidor

### 📋 LO QUE NECESITAMOS

El equipo frontend hace dos llamadas separadas: una a `/users/1` y otra a `/status`. Para mejorar el rendimiento, necesitamos un nuevo endpoint `/user-info/{id}` que devuelva ambos datos en una sola respuesta.

### 🏛️ ARQUITECTURA CORRECTA

**⚠️ IMPORTANTE**: En este ejercicio, el **Handler orquesta** múltiples casos de uso, NO creamos un UseCase que llame a otros.

**✅ CORRECTO (Handler orquesta)**:
```
Handler → Llama GetUserUsecase
       → Llama GetStatusUsecase  
       → Combina resultados
       → Devuelve JSON
```

**❌ INCORRECTO (UseCase orquesta otros UseCases)**:
```
Handler → GetUserInfoUsecase → GetUserUsecase
                             → GetStatusUsecase
```

**¿Por qué?** El Handler es responsable de la orquestación HTTP, los UseCases deben ser independientes y reutilizables.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. Ve a: `internal/infrastructure/http/router.go`
2. Comenta temporalmente la nueva ruta:
   ```go
   // r.Get("/user-info/{id}", userInfoHandler.GetByID)
   ```

**PASO 2: Crear el Handler (AQUÍ VA LA ORQUESTACIÓN)**

3. En `internal/adapter/http/handler/`
4. Crea archivo: `user_info_handler.go`
5. Crea `UserInfoHandler` que **recibe AMBOS casos de uso**:
   ```go
   type UserInfoHandler struct {
       getUserUC   *userUsecase.GetUserUsecase
       getStatusUC *statusUsecase.GetStatusUsecase
   }
   ```

6. Crea el método `GetByID` que **orquesta**:
   - Extrae y valida el parámetro `id` de la URL
   - **Llama directamente** a `h.getUserUC.Execute(id)`
   - **Llama directamente** a `h.getStatusUC.Execute()`
   - **Combina resultados** en una estructura anónima:
     ```go
     response := struct {
         User         *user.User     `json:"user"`
         ServerStatus *status.Status `json:"server_status"`
     }{
         User:         userData,
         ServerStatus: statusData,
     }
     ```
   - Devuelve la respuesta en JSON

**PASO 3: Conectar en main.go**

7. Ve a: `cmd/app/main.go`
8. En la sección de handlers, crea el handler pasándole **ambos casos de uso**:
   ```go
   userInfoHandler := handler.NewUserInfoHandler(getUserUsecase, getStatusUsecase)
   ```

**PASO 4: Actualizar el Router**

9. Ve a: `internal/infrastructure/http/router.go`
10. Agrega el parámetro `userInfoHandler` a la función `SetupRouter`
11. Activa la ruta: `r.Get("/user-info/{id}", userInfoHandler.GetByID)`

**PASO 5: Actualizar main.go para pasar el handler**

12. En `cmd/app/main.go`, cuando llamas a `SetupRouter`, pasa el nuevo handler
13. Reinicia el servidor

**PASO 6: Prueba**

14. Ejecuta: `curl http://localhost:8080/user-info/1`
15. Debe mostrar usuario + status en una respuesta

### ✅ RESULTADO ESPERADO

```json
{
  "user": { "id": 1, "name": "...", ... },
  "server_status": { "message": "...", "version": "1.1.0", ... }
}
```

### 💡 LO QUE APRENDISTE

**Conceptos Clave**:
- ✅ El **Handler orquesta** llamadas a múltiples casos de uso
- ✅ Los **UseCases permanecen independientes** y reutilizables
- ✅ No creamos UseCases que solo llamen a otros UseCases
- ✅ La **composición de datos** ocurre en el Handler (capa de adaptador HTTP)
- ✅ Cada UseCase tiene una sola responsabilidad

**Ventajas de este enfoque**:
- Los UseCases son 100% reutilizables desde CLI, Workers, Tests, otros Handlers
- No hay acoplamiento entre casos de uso
- Más fácil de testear cada pieza independientemente
- Respeta el principio de responsabilidad única

---

## EJERCICIO 13 - Validar que el nombre no esté vacío (preparación)

### 📋 LO QUE NECESITAMOS

Preparación para futuros endpoints: Si en el futuro agregamos un endpoint que reciba un nombre, necesitamos asegurarnos de que no esté vacío. Por ahora, solo agrega la función de validación.

### 🎯 INSTRUCCIONES

1. **Crea un paquete de validaciones**
   - Carpeta: `internal/domain/validation/`
   - Archivo: `text_validation.go`

2. **Crea la función**
   - Función: `ValidateNotEmpty(text string) error`
   - Si el texto está vacío, devuelve error: "el campo no puede estar vacío"
   - Si no está vacío, devuelve nil (sin error)

3. **Prueba manual** (opcional)
   - Crea un test en `test/domain/validation/`
   - Verifica que rechaza textos vacíos

### ✅ RESULTADO ESPERADO

Tienes una función reutilizable para validar textos en futuros endpoints.

### 💡 LO QUE HICISTE

Creaste una utilidad que se puede usar en múltiples lugares.

---

## EJERCICIO 14 - Crear módulo completo de productos

### 📋 LO QUE NECESITAMOS

El cliente quiere agregar productos a la aplicación. Necesitamos crear todo un módulo nuevo que se conecte a `https://fakestoreapi.com`:

- Endpoint para obtener un producto: `/products/{id}`
- Endpoint para listar productos: `/products`

Un producto tiene: ID, Título, Precio, Descripción, Categoría.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar las rutas**

1. Ve a: `internal/infrastructure/http/router.go`
2. Comenta temporalmente:
   ```go
   // r.Get("/products", productHandler.List)
   // r.Get("/products/{id}", productHandler.GetByID)
   ```

**PASO 2: Crear el Handler**

3. En `internal/adapter/http/handler/`
4. Crea: `product_handler.go`
5. Crea `ProductHandler` con dos casos de uso (get y list)
6. Implementa dos métodos:
   - `GetByID`: extrae ID, valida, llama caso de uso
   - `List`: llama caso de uso y devuelve lista

**PASO 3: Crear el modelo de Dominio (datos que manejaremos)**

7. Crea carpeta: `internal/domain/product/`
8. Crea archivo: `product.go`
9. Define la estructura `Product` con estos campos:
   - ID: tipo `int`, tag JSON `"id"`
   - Title: tipo `string`, tag JSON `"title"`
   - Price: tipo `float64`, tag JSON `"price"`
   - Description: tipo `string`, tag JSON `"description"`
   - Category: tipo `string`, tag JSON `"category"`
   
   💡 **Tip**: Revisa cómo está definido `User` en `internal/domain/user/user.go` para ver el formato

**PASO 4: Definir el contrato del Repositorio (lo que necesitamos obtener)**

10. En la misma carpeta `internal/domain/product/`
11. Crea archivo: `repository.go`
12. Define una interfaz llamada `Repository` con dos métodos:
    - `FindByID(id int) (*Product, error)` - para obtener un producto
    - `FindAll() ([]*Product, error)` - para obtener todos los productos
    
    💡 **Tip**: Revisa `internal/domain/user/repository.go` para ver el patrón

**PASO 5: Crear los Casos de Uso (la lógica de negocio)**

12. Crea carpeta: `internal/usecase/product/`
13. Crea archivo: `get_product.go`
    - Recibe el repositorio en el constructor
    - Valida que el ID esté entre 1 y 20
    - Llama a `repository.FindByID(id)`
14. Crea archivo: `list_products.go`
    - Recibe el repositorio en el constructor
    - Llama a `repository.FindAll()`

**PASO 6: Implementar el Repositorio (obtener datos de la API externa)**

15. Ve a: `internal/adapter/repository/`
16. Crea archivo: `product_api_repository.go`
17. Implementa ambos métodos:
    - `FindByID`: GET a `https://fakestoreapi.com/products/{id}`
    - `FindAll`: GET a `https://fakestoreapi.com/products`
    - Maneja errores apropiadamente (404, 500, etc.)

**PASO 7: Conectar todo en el Router (montar la aplicación)**

18. Vuelve a `router.go`
19. Crea la instancia del repositorio de productos
20. Crea las instancias de ambos casos de uso (pásales el repositorio)
21. Crea la instancia del handler (pásale los casos de uso)
22. Descomenta y activa las rutas

**PASO 8: Prueba**

22. Reinicia el servidor
23. Prueba:
    - `curl http://localhost:8080/products` → Lista
    - `curl http://localhost:8080/products/1` → Un producto
    - `curl http://localhost:8080/products/999` → Error de validación

### ✅ RESULTADO ESPERADO

- `/products` → Lista de productos (20 productos)
- `/products/1` → Un producto específico
- `/products/21` → Error: "el ID debe estar entre 1 y 20"
- `/products/0` → Error: "el ID debe ser mayor que 0"

### 💡 LO QUE HICISTE

Creaste un módulo completo desde cero siguiendo el flujo correcto, ahora puedes replicar esta estructura para cualquier recurso nuevo.

---

## EJERCICIO 15 - Endpoint para buscar usuarios por email

### 🏢 CONTEXTO DE NEGOCIO

**Problema del cliente**:
El equipo de soporte técnico recibe llamadas de usuarios que olvidaron su nombre de usuario. Solo recuerdan su email. El soporte necesita buscar rápidamente el perfil del usuario usando su email para poder ayudarlos.

**Por qué es importante**:
- **Soporte**: Reduce tiempo de atención de 5 minutos a 30 segundos
- **Experiencia de usuario**: Usuario frustrado recibe ayuda inmediata
- **Costos**: Menos tiempo = menos costo operativo

**Caso de uso real**:
1. Usuario llama: "No puedo entrar, olvidé mi nombre de usuario"
2. Soporte pregunta: "¿Cuál es tu email registrado?"
3. Soporte busca: `GET /users/search?email=usuario@example.com`
4. Soporte encuentra el perfil y lo ayuda inmediatamente

---

### 🏛️ POR QUÉ CLEAN ARCHITECTURE EN ESTE EJERCICIO

**Separación en capas** (de afuera hacia adentro):

1. **Handler** (Adaptador HTTP)
   - **Qué hace**: Recibe petición HTTP, extrae parámetros, devuelve respuesta HTTP
   - **Por qué separado**: Si mañana cambiamos a gRPC o GraphQL, solo cambiamos esta capa
   - **Beneficio**: La lógica de negocio no se contamina con detalles de HTTP

2. **Use Case** (Lógica de Negocio)
   - **Qué hace**: Implementa la regla "buscar usuario por email"
   - **Por qué separado**: Es reutilizable desde HTTP, CLI, Workers, Tests
   - **Beneficio**: Si lanzamos app móvil, reutilizamos este mismo caso de uso

3. **Repository** (Puerto de Datos)
   - **Qué hace**: Obtiene usuarios (ahora de API, mañana de BD)
   - **Por qué separado**: Podemos cambiar de API a PostgreSQL sin tocar el Use Case
   - **Beneficio**: Intercambiable y testeable con mocks

**Ejemplo real de beneficio**:
Si en 6 meses migramos de `jsonplaceholder` a nuestra BD:
- ✅ Handler: No cambia
- ✅ Use Case: No cambia
- ❌ Repository: Solo este archivo cambia (de HTTP a SQL)

Esfuerzo: 1 archivo vs reescribir todo.

---

### 📋 LO QUE NECESITAMOS

**Endpoint**: `GET /users/search?email=ejemplo@dominio.com`

**Entrada**:
- Query parameter `email` (string, obligatorio)

**Salida exitosa** (200):
```json
{
  "id": 1,
  "name": "Leanne Graham",
  "email": "Sincere@april.biz",
  ...
}
```

**Errores**:
- 400: Email vacío → `{"error": "el email es requerido"}`
- 404: No encontrado → `{"error": "usuario no encontrado"}`

---

### ✅ VALIDACIONES OBLIGATORIAS

#### 1. Validar email no vacío
- **Dónde**: Handler (entrada) + Use Case (defensa)
- **Por qué**: Evitar búsquedas innecesarias
- **Si no se hace**: Llamada inútil al repositorio, error confuso
- **Código**:
  ```go
  if email == "" {
      http.Error(w, `{"error":"email requerido"}`, 400)
      return
  }
  ```

#### 2. Retornar 404 si no existe
- **Dónde**: Use Case
- **Por qué**: Semántica HTTP correcta
- **Si no se hace**: Cliente no sabrá si es error de servidor o simplemente no existe
- **Código**:
  ```go
  if usuarioEncontrado == nil {
      return nil, errors.New("usuario no encontrado")
  }
  ```

---

### ⚠️ MALAS PRÁCTICAS A EVITAR

#### ❌ MAL: Lógica de negocio en el Handler

```go
// ❌ NO HACER
func (h *UserHandler) SearchByEmail(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    // ❌ Buscar directamente aquí
    users, _ := h.repository.FindAll()
    for _, user := range users {
        if user.Email == email {
            json.NewEncoder(w).Encode(user)
            return
        }
    }
}
```

**Problema**: No es reutilizable, no es testeable, mezcla responsabilidades.

#### ✅ BIEN: Handler delgado

```go
// ✅ SÍ HACER
func (h *UserHandler) SearchByEmail(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    if email == "" {
        http.Error(w, `{"error":"email requerido"}`, 400)
        return
    }
    // ✅ Delegar al caso de uso
    user, err := h.searchUC.Execute(email)
    if err != nil {
        http.Error(w, `{"error":"no encontrado"}`, 404)
        return
    }
    json.NewEncoder(w).Encode(user)
}
```

#### ❌ MAL: Ignorar errores

```go
// ❌ NO HACER
user, _ := h.searchUC.Execute(email)  // Ignora error
json.NewEncoder(w).Encode(user)       // ¡user puede ser nil! → PANIC
```

**Consecuencia**: Aplicación se cae, usuario recibe "Internal Server Error".

#### ✅ BIEN: Manejar errores

```go
// ✅ SÍ HACER
user, err := h.searchUC.Execute(email)
if err != nil {
    http.Error(w, `{"error":"no encontrado"}`, 404)
    return
}
json.NewEncoder(w).Encode(user)
```

---

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. Ve a: `internal/infrastructure/http/router.go`
2. Comenta temporalmente (importante: debe ir ANTES de `/users/{id}`):
   ```go
   // r.Get("/users/search", userHandler.SearchByEmail)
   ```

**PASO 2: Crear el método en el Handler**

3. Ve a: `internal/adapter/http/handler/user_handler.go`
4. Agrega el nuevo caso de uso a la estructura del handler
5. Crea el método `SearchByEmail`:
   - Obtiene el parámetro `email` de la query: `r.URL.Query().Get("email")`
   - Valida que el email no esté vacío
   - Llama al caso de uso
   - Devuelve el usuario en JSON

**PASO 3: Crear el Caso de Uso**

6. En `internal/usecase/user/`
7. Crea: `search_user_by_email.go`
8. Implementa `SearchUserByEmailUsecase`:
   - Valida que el email no esté vacío
   - Llama al repositorio para obtener todos los usuarios
   - Busca en la lista el usuario con ese email
   - Si no encuentra, devuelve error "usuario no encontrado"

**PASO 4: Conectar en el Router**

9. Vuelve a `router.go`
10. Crea el caso de uso (usa el mismo repositorio de usuarios)
11. Pásalo al handler al crearlo
12. Activa la ruta (asegúrate que esté ANTES de `/users/{id}`)

**PASO 5: Prueba**

13. Reinicia el servidor
14. Prueba: `curl "http://localhost:8080/users/search?email=Sincere@april.biz"`
15. Debe devolver el usuario con ID 1

### ✅ RESULTADO ESPERADO

- `/users/search?email=Sincere@april.biz` → Usuario encontrado
- `/users/search?email=noexiste@test.com` → Error: "usuario no encontrado"
- `/users/search` → Error: "email es requerido"

### 💡 LO QUE APRENDISTE

#### Conceptos Técnicos:
- ✅ Query parameters en Go: `r.URL.Query().Get("param")`
- ✅ Filtrado lineal en slices
- ✅ Comparación case-insensitive: `strings.EqualFold()`
- ✅ Códigos HTTP apropiados: 400 vs 404

#### Conceptos Arquitectónicos:
- ✅ **Separación de responsabilidades**: Handler (HTTP) ≠ Use Case (negocio)
- ✅ **Reutilización**: El Use Case funciona desde HTTP, CLI, tests
- ✅ **Defensa en profundidad**: Validar en Handler Y Use Case
- ✅ **Inversión de dependencias**: Use Case no conoce HTTP

#### Habilidades de Negocio:
- ✅ Entender requisitos desde perspectiva del usuario final
- ✅ Identificar validaciones y su impacto en experiencia
- ✅ Documentar casos de uso reales

**Este mismo patrón lo usarás para**: Buscar productos por nombre, buscar pedidos por tracking, cualquier búsqueda por criterio específico.

---

## EJERCICIO 16 - Endpoint para obtener solo el email de un usuario

### 📋 LO QUE NECESITAMOS

A veces solo necesitamos el email de un usuario, no toda su información. Crea un endpoint `/users/{id}/email` que devuelva únicamente el email.

**NOTA**: Este ejercicio reutiliza el caso de uso existente de GetUser, solo crea una vista diferente de los datos.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/users/{id}/email", userHandler.GetEmail)
   ```

**PASO 2: Crear estructura de respuesta (Dominio)**

2. En `internal/domain/user/`
3. Crea archivo: `user_email.go`
4. Define estructura simple:
   ```go
   type UserEmail struct {
       Email string `json:"email"`
   }
   ```

**PASO 3: Crear el método en el Handler**

5. En `user_handler.go`
6. Crea método `GetEmail`:
   - Extrae el ID del parámetro de ruta
   - Convierte a int y valida
   - Llama al caso de uso **existente** GetUser
   - Extrae solo el email y créalo en una estructura UserEmail
   - Devuelve en JSON

**PASO 4: Conectar en el Router**

7. En `router.go`
8. Registra la ruta: `r.Get("/users/{id}/email", userHandler.GetEmail)`
9. **NO necesitas crear nuevo caso de uso**, reutiliza el existente

**PASO 5: Prueba**

10. `curl http://localhost:8080/users/1/email`
11. Debe mostrar: `{"email":"Sincere@april.biz"}`

### ✅ RESULTADO ESPERADO

- `/users/1/email` → Solo el email
- `/users/999/email` → Error: "el ID debe estar entre 1 y 10"

### 💡 LO QUE HICISTE

Aprendiste a crear endpoints especializados que devuelven información parcial reutilizando casos de uso existentes.

---

## EJERCICIO 17 - Endpoint para buscar productos por categoría

### 📋 LO QUE NECESITAMOS

Queremos filtrar productos por categoría. Crea `/products/category/{category}` que devuelva todos los productos de esa categoría.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/products/category/{category}", productHandler.GetByCategory)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Agrega el nuevo caso de uso
4. Crea método `GetByCategory`:
   - Extrae el parámetro `category`
   - Llama al caso de uso
   - Devuelve la lista en JSON

**PASO 3: Crear el Caso de Uso**

5. En `internal/usecase/product/`
6. Crea: `get_products_by_category.go`
7. Implementa `GetProductsByCategoryUsecase`:
   - Valida que la categoría no esté vacía
   - Obtiene todos los productos del repositorio
   - Filtra los que coincidan con la categoría (case-insensitive)
   - Devuelve la lista filtrada

**PASO 4: Conectar en el Router**

8. Crea el caso de uso
9. Pásalo al handler
10. Activa la ruta

**PASO 5: Prueba**

11. `curl http://localhost:8080/products/category/electronics`
12. Debe devolver solo productos electrónicos

### ✅ RESULTADO ESPERADO

- `/products/category/electronics` → Lista de productos electrónicos
- `/products/category/jewelery` → Lista de joyas
- `/products/category/noexiste` → Lista vacía `[]`

### 💡 LO QUE HICISTE

Implementaste filtrado de datos con parámetros de ruta.

---

## EJERCICIO 18 - Endpoint para productos con precio mayor a X

### 📋 LO QUE NECESITAMOS

El cliente quiere ver solo productos caros. Crea `/products/price-above?min=100` que devuelva productos con precio mayor al especificado.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/price-above", productHandler.GetAbovePrice)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Agrega el caso de uso
4. Crea `GetAbovePrice`:
   - Obtiene parámetro `min` de la query
   - Convierte a float64
   - Llama al caso de uso
   - Devuelve lista en JSON

**PASO 3: Crear el Caso de Uso**

5. Crea: `internal/usecase/product/get_products_above_price.go`
6. Implementa `GetProductsAbovePriceUsecase`:
   - Valida que minPrice sea mayor que 0
   - Obtiene todos los productos
   - Filtra los que tengan precio >= minPrice
   - Devuelve la lista

**PASO 4: Conectar en el Router**

7. Crea el caso de uso
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl "http://localhost:8080/products/price-above?min=100"`
11. Debe mostrar solo productos caros

### ✅ RESULTADO ESPERADO

- `/products/price-above?min=100` → Productos >= $100
- `/products/price-above?min=1000` → Productos muy caros (o lista vacía)
- `/products/price-above` → Error: "precio mínimo es requerido"

### 💡 LO QUE HICISTE

Combinaste query parameters con filtrado numérico.

---

## EJERCICIO 19 - Endpoint para filtrar usuarios por dominio de email

### 📋 LO QUE NECESITAMOS

El equipo de marketing quiere saber cuántos usuarios tenemos de cada dominio de email (gmail.com, outlook.com, etc.). Necesitamos `/users/by-domain?domain=gmail.com` que devuelva todos los usuarios de ese dominio.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/users/{id}`):
   ```go
   // r.Get("/users/by-domain", userHandler.GetByDomain)
   ```

**PASO 2: Crear el método en el Handler**

2. En `user_handler.go`
3. Crea método `GetByDomain`:
   - Obtiene parámetro `domain` de la query
   - Valida que no esté vacío
   - Llama al caso de uso
   - Devuelve lista de usuarios en JSON

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/user/`
5. Crea: `list_users_by_domain.go`
6. Implementa `ListUsersByDomainUsecase`:
   - Valida que el dominio no esté vacío
   - Obtiene todos los usuarios del repositorio
   - Filtra usuarios cuyo email termine con `@dominio`
   - Tip: Usa `strings.HasSuffix(user.Email, "@"+domain)`
   - Devuelve la lista filtrada

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl "http://localhost:8080/users/by-domain?domain=biz"` → Usuarios con email @biz
11. `curl "http://localhost:8080/users/by-domain?domain=april.biz"` → Solo 1 usuario
12. `curl "http://localhost:8080/users/by-domain?domain=noexiste.com"` → Lista vacía []

### ✅ RESULTADO ESPERADO

- `/users/by-domain?domain=biz` → Lista de usuarios con emails `@*.biz`
- `/users/by-domain?domain=noexiste.com` → `[]` (lista vacía, no error)
- `/users/by-domain` → Error: "dominio es requerido"

### 💡 LO QUE HICISTE

Practicaste filtrado con `strings.HasSuffix()` y manejo de resultados vacíos (no es error, es una lista válida sin elementos).

---

## EJERCICIO 20 - Endpoint para listar solo títulos de productos

### 📋 LO QUE NECESITAMOS

Para un dropdown en el frontend, necesitan `/products/titles` que devuelva una lista simple con solo los títulos de los productos, sin toda la información completa.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/titles", productHandler.GetTitles)
   ```

**PASO 2: Crear estructura de respuesta**

2. En `internal/domain/product/`
3. Crea archivo: `product_title.go`
4. Define `ProductTitle` con dos campos: `ID` (int) y `Title` (string)

**PASO 3: Crear el método en el Handler**

5. En `product_handler.go`
6. Crea método `GetTitles`:
   - Llama al caso de uso existente ListProducts
   - Recorre todos los productos
   - Crea una lista de ProductTitle solo con ID y título
   - Devuelve la lista en JSON

**PASO 4: Conectar en el Router**

7. En `router.go`
8. Registra la ruta (reutiliza el handler y caso de uso de ListProducts)

**PASO 5: Prueba**

9. `curl http://localhost:8080/products/titles`
10. Debe mostrar lista simplificada

### ✅ RESULTADO ESPERADO

```json
[
  {"id": 1, "title": "Fjallraven - Foldsack No. 1 Backpack"},
  {"id": 2, "title": "Mens Casual Premium Slim Fit T-Shirts"},
  ...
]
```

### 💡 LO QUE HICISTE

Transformaste datos existentes para crear una vista simplificada, útil para optimizar respuestas del API.

---

## EJERCICIO 21 - Endpoint para contar productos por categoría

### 📋 LO QUE NECESITAMOS

El cliente quiere saber cuántos productos hay en cada categoría sin tener que traer todos los productos. Crea `/products/count-by-category` que devuelva un resumen.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/products/count-by-category", productHandler.CountByCategory)
   ```

**PASO 2: Crear estructura de respuesta**

2. En `internal/domain/product/`
3. Crea archivo: `category_count.go`
4. Define `CategoryCount` con: `Category` (string) y `Count` (int)

**PASO 3: Crear el método en el Handler**

5. En `product_handler.go`
6. Crea método `CountByCategory`:
   - Llama al caso de uso ListProducts
   - Crea un mapa para contar: `map[string]int`
   - Recorre productos y cuenta por categoría
   - Convierte el mapa a una lista de CategoryCount
   - Devuelve en JSON

**PASO 4: Conectar en el Router**

7. Registra la ruta (reutiliza caso de uso existente)

**PASO 5: Prueba**

8. `curl http://localhost:8080/products/count-by-category`

### ✅ RESULTADO ESPERADO

```json
[
  {"category": "electronics", "count": 6},
  {"category": "jewelery", "count": 4},
  {"category": "men's clothing", "count": 4},
  {"category": "women's clothing", "count": 6}
]
```

### 💡 LO QUE HICISTE

Agregaste lógica de agregación simple en el handler, procesando datos en memoria para crear resúmenes.

---

## EJERCICIO 22 - Endpoint para buscar productos por título

### 📋 LO QUE NECESITAMOS

Necesitamos buscar productos por texto en el título. Crea `/products/search?query=shirt` que devuelva todos los productos cuyo título contenga esa palabra (sin distinguir mayúsculas).

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/search", productHandler.SearchByTitle)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `SearchByTitle`:
   - Obtiene parámetro `query` de la URL
   - Valida que no esté vacío
   - Llama al caso de uso
   - Devuelve los resultados en JSON

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `search_products_by_title.go`
6. Implementa `SearchProductsByTitleUsecase`:
   - Valida que query no esté vacío
   - Obtiene todos los productos del repositorio
   - Filtra los que contengan el texto en el título (usar `strings.Contains` y `strings.ToLower`)
   - Devuelve la lista filtrada

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio de productos)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl "http://localhost:8080/products/search?query=shirt"`
11. Debe mostrar productos con "shirt" en el título

### ✅ RESULTADO ESPERADO

- `/products/search?query=shirt` → Productos con "shirt" en el título
- `/products/search?query=jacket` → Productos con "jacket"
- `/products/search` → Error: "query es requerida"
- `/products/search?query=xyz` → Lista vacía `[]`

### 💡 LO QUE HICISTE

Implementaste búsqueda de texto básica, filtrando información en memoria de forma case-insensitive.

---

## EJERCICIO 23 - Endpoint para ordenar productos por precio ascendente

### 📋 LO QUE NECESITAMOS

Para el frontend de la tienda, necesitan mostrar productos del más barato al más caro. Crea `/products/by-price` que devuelva todos los productos ordenados por precio ascendente.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/by-price", productHandler.GetByPrice)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetByPrice`:
   - Llama al caso de uso (no necesita parámetros)
   - Devuelve lista ordenada en JSON

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `list_products_by_price.go`
6. Implementa `ListProductsByPriceUsecase`:
   - Obtiene todos los productos del repositorio
   - Usa `sort.Slice(products, func(i, j int) bool { ... })`
   - Compara: `return products[i].Price < products[j].Price`
   - Tip: Importa el package `sort`
   - Devuelve la lista ordenada

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl http://localhost:8080/products/by-price`
11. Verifica que el primero sea el más barato y el último el más caro

### ✅ RESULTADO ESPERADO

Lista de productos ordenada del menor al mayor precio.

### 💡 LO QUE HICISTE

Aprendiste a usar `sort.Slice()` para ordenar estructuras personalizadas por un campo específico.

---

## EJERCICIO 24 - Endpoint para listar productos ordenados por precio

### 📋 LO QUE NECESITAMOS

Crea `/products/sorted-by-price` que devuelva todos los productos ordenados de menor a mayor precio.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/sorted-by-price", productHandler.GetSortedByPrice)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetSortedByPrice`:
   - Llama al caso de uso ListProducts
   - Ordena los productos por precio usando `sort.Slice`
   - Devuelve la lista ordenada

**PASO 3: Conectar y Probar**

4. Registra la ruta (reutiliza caso de uso ListProducts)
5. `curl http://localhost:8080/products/sorted-by-price`

### ✅ RESULTADO ESPERADO

Lista de productos ordenada del más barato al más caro.

### 💡 LO QUE HICISTE

Aplicaste ordenamiento a una colección de datos antes de devolverla.

---

## EJERCICIO 25 - Endpoint para combinar filtros de categoría y precio

### 📋 LO QUE NECESITAMOS

El frontend de la tienda necesita filtrar productos por categoría Y por precio mínimo al mismo tiempo. Crea `/products/filter?category=electronics&minPrice=100` que combine ambos filtros.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/filter", productHandler.GetFiltered)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetFiltered`:
   - Obtiene `category` (string) de la query
   - Obtiene `minPrice` (string) de la query, conviértelo a float64 con `strconv.ParseFloat()`
   - Llama al caso de uso con ambos parámetros
   - Devuelve lista filtrada

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `filter_products.go`
6. Implementa `FilterProductsUsecase`:
   - Recibe category (string) y minPrice (float64)
   - Obtiene todos los productos
   - Filtra en dos pasos:
     - Primero por categoría (si category no está vacío)
     - Luego por precio >= minPrice (si minPrice > 0)
   - Devuelve productos que cumplen AMBOS criterios

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl "http://localhost:8080/products/filter?category=electronics&minPrice=100"`
11. Verifica que todos sean electrónicos Y tengan precio >= 100

### ✅ RESULTADO ESPERADO

Lista de productos que cumplen ambos filtros simultáneamente.

### 💡 LO QUE HICISTE

Practicaste combinar múltiples filtros en secuencia, un patrón común en búsquedas avanzadas.

---

## EJERCICIO 26 - Endpoint para encontrar el producto más barato

### 📋 LO QUE NECESITAMOS

Para promociones, necesitamos saber cuál es el producto más barato de la tienda. Crea `/products/cheapest` que devuelva el producto con el precio más bajo.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/products/cheapest", productHandler.GetCheapest)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetCheapest`:
   - Llama al caso de uso
   - Devuelve el producto en JSON

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `get_cheapest_product.go`
6. Implementa `GetCheapestProductUsecase`:
   - Obtiene todos los productos
   - Inicializa una variable `cheapest` con el primer producto
   - Recorre todos los productos
   - Si encuentra uno con precio menor, actualiza `cheapest`
   - Devuelve el producto más barato

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl http://localhost:8080/products/cheapest`
11. Verifica que sea el producto con el precio más bajo

### ✅ RESULTADO ESPERADO

El producto con el precio más bajo de todos.

### 💡 LO QUE HICISTE

Aprendiste a buscar el elemento con el valor mínimo en una colección.

---

## EJERCICIO 27 - Endpoint para filtrar productos por rango de precio

### 📋 LO QUE NECESITAMOS

Para filtros avanzados, necesitamos buscar productos dentro de un rango de precio. Crea `/products/in-range?min=50&max=150` que devuelva productos entre esos precios (inclusive).

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/in-range", productHandler.GetInRange)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetInRange`:
   - Obtiene `min` y `max` de la query
   - Convierte ambos a float64
   - Valida que min < max
   - Llama al caso de uso
   - Devuelve lista filtrada

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `list_products_in_range.go`
6. Implementa `ListProductsInRangeUsecase`:
   - Valida que min <= max
   - Obtiene todos los productos
   - Filtra productos donde: `price >= min AND price <= max`
   - Devuelve lista filtrada

**PASO 4: Conectar en el Router**

7. Crea el caso de uso (usa el mismo repositorio)
8. Pásalo al handler
9. Activa la ruta

**PASO 5: Prueba**

10. `curl "http://localhost:8080/products/in-range?min=50&max=150"`
11. Verifica que todos estén entre 50 y 150

### ✅ RESULTADO ESPERADO

Lista de productos con precio entre min y max (inclusive).

### 💡 LO QUE HICISTE

Practicaste filtrado con rangos numéricos y validaciones de lógica de negocio (min < max).

---

## EJERCICIO 28 - Endpoint para obtener ciudad de un usuario

### 📋 LO QUE NECESITAMOS

Para estadísticas geográficas, necesitamos extraer la ciudad donde vive un usuario. Crea `/users/{id}/city` que devuelva solo la ciudad.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/users/{id}/city", userHandler.GetCity)
   ```

**PASO 2: Crear estructura de respuesta**

2. En `internal/domain/user/`
3. Crea archivo: `user_city.go`
4. Define estructura `UserCity` con un campo: `City string`

**PASO 3: Crear el método en el Handler**

5. En `user_handler.go`
6. Crea método `GetCity`:
   - Extrae el ID del path parameter
   - Valida y convierte a int
   - Llama al caso de uso GetUser (reutiliza el existente)
   - Accede a la ciudad: `user.Address.City`
   - Crea estructura UserCity con esa ciudad
   - Devuelve en JSON

**PASO 4: Conectar en el Router**

7. Registra la ruta (reutiliza el handler y caso de uso existente)

**PASO 5: Prueba**

8. `curl http://localhost:8080/users/1/city`
9. Debe mostrar solo: `{"city": "Gwenborough"}`

### ✅ RESULTADO ESPERADO

- `/users/1/city` → `{"city": "Gwenborough"}`
- `/users/999/city` → Error 404

### 💡 LO QUE HICISTE

Aprendiste a acceder a datos anidados (`user.Address.City`) y crear respuestas parciales reutilizando lógica existente.

---

## EJERCICIO 29 - Endpoint para filtrar productos por rango de precio

### 📋 LO QUE NECESITAMOS

Crea `/products/price-range?min=50&max=200` que devuelva productos dentro de ese rango.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/products/price-range", productHandler.GetByPriceRange)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Crea método `GetByPriceRange`:
   - Obtiene parámetros `min` y `max` de la query
   - Convierte a float64
   - Valida que min < max
   - Llama al caso de uso

**PASO 3: Crear el Caso de Uso**

4. En `internal/usecase/product/`
5. Crea: `get_products_by_price_range.go`
6. Implementa el filtrado:
   - Obtiene todos los productos
   - Filtra los que estén entre min y max
   - Devuelve la lista

**PASO 4: Conectar y Probar**

7. `curl "http://localhost:8080/products/price-range?min=50&max=200"`

### ✅ RESULTADO ESPERADO

Lista de productos entre $50 y $200.

### 💡 LO QUE HICISTE

Implementaste filtrado con dos parámetros simultáneos.

---

## EJERCICIO 30 - Endpoint para listar categorías únicas

### 📋 LO QUE NECESITAMOS

Crea `/products/categories` que devuelva una lista de todas las categorías únicas disponibles.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/products/categories", productHandler.GetCategories)
   ```

**PASO 2: Crear estructura de respuesta**

2. En `internal/domain/product/`
3. Crea: `category_list.go`
4. Define `CategoryList` con: `Categories []string`

**PASO 3: Crear el método en el Handler**

5. En `product_handler.go`
6. Crea método `GetCategories`:
   - Llama a ListProducts
   - Crea un map para almacenar categorías únicas
   - Recorre productos y agrega categorías al map
   - Convierte map a slice
   - Devuelve la lista

**PASO 4: Conectar y Probar**

7. `curl http://localhost:8080/products/categories`

### ✅ RESULTADO ESPERADO

```json
{
  "categories": [
    "electronics",
    "jewelery",
    "men's clothing",
    "women's clothing"
  ]
}
```

### 💡 LO QUE HICISTE

Extrajiste valores únicos de una colección usando un mapa.

---

## 🔥 NIVEL AVANZADO - Endpoints Complejos

A partir de aquí, los ejercicios combinan múltiples fuentes de datos y requieren más planificación.

---

## EJERCICIO 31 - Endpoint de estadísticas de usuarios y productos

### 📋 LO QUE NECESITAMOS

Para el dashboard, necesitamos `/stats` que muestre:
- Total de usuarios disponibles
- Total de productos disponibles
- Categorías de productos únicas
- Tiempo de respuesta promedio estimado

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: El **Handler orquesta** múltiples casos de uso y procesa/agrega los datos.

```
StatsHandler → Llama ListUsersUsecase → Cuenta total
             → Llama ListProductsUsecase → Cuenta total + extrae categorías
             → Construye respuesta agregada
             → Devuelve JSON
```

**NO** creamos un `GetStatsUsecase` que llame a otros UseCases. Los cálculos simples (conteos, agregaciones) se hacen en el Handler.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/stats", statsHandler.Get)
   ```

**PASO 2: Crear el Handler (ORQUESTACIÓN Y AGREGACIÓN)**

2. Crea: `internal/adapter/http/handler/stats_handler.go`
3. Define `StatsHandler` que recibe **ambos casos de uso**:
   ```go
   type StatsHandler struct {
       listUsersUC    *userUsecase.ListUsersUsecase
       listProductsUC *productUsecase.ListProductsUsecase
   }
   ```

4. Implementa el método `Get`:
   - Llama a `h.listUsersUC.Execute()` y cuenta: `len(users)`
   - Llama a `h.listProductsUC.Execute()` y cuenta: `len(products)`
   - Extrae categorías únicas usando un `map[string]bool`
   - Construye la respuesta con estructura anónima:
     ```go
     response := struct {
         TotalUsers    int      `json:"total_users"`
         TotalProducts int      `json:"total_products"`
         Categories    []string `json:"categories"`
         ResponseTime  string   `json:"response_time"`
     }{
         TotalUsers:    len(users),
         TotalProducts: len(products),
         Categories:    categories,
         ResponseTime:  "~500ms",
     }
     ```
   - Devuelve JSON

**PASO 3: Conectar en main.go**

5. Crea el handler pasándole ambos casos de uso:
   ```go
   statsHandler := handler.NewStatsHandler(listUsersUsecase, listProductsUsecase)
   ```

**PASO 4: Actualizar el Router**

6. Agrega `statsHandler` como parámetro a `SetupRouter`
7. Activa la ruta: `r.Get("/stats", statsHandler.Get)`

**PASO 5: Prueba**

8. Reinicia el servidor
9. `curl http://localhost:8080/stats`

### ✅ RESULTADO ESPERADO

```json
{
  "total_users": 10,
  "total_products": 20,
  "categories": ["electronics", "jewelery", "men's clothing", "women's clothing"],
  "response_time": "~500ms"
}
```

### 💡 LO QUE APRENDISTE

**Conceptos Clave**:
- ✅ Agregaciones simples (conteos, filtros) van en el **Handler**
- ✅ UseCases devuelven datos crudos, el Handler los transforma
- ✅ No necesitas un UseCase para cada endpoint si solo combinas datos
- ✅ El Handler puede hacer lógica de presentación (formateo, agregación simple)

---

## EJERCICIO 32 - Endpoint de perfil completo de usuario

### 📋 LO QUE NECESITAMOS

Crea un endpoint `/users/{id}/profile` que devuelva:
- Información completa del usuario
- Total de usuarios en el sistema (para mostrar "Usuario X de Y")
- Un mensaje personalizado: "Perfil de [nombre]"

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: Handler orquesta y formatea.

```
UserProfileHandler → Llama GetUserUsecase(id)
                   → Llama ListUsersUsecase → cuenta len()
                   → Construye mensaje personalizado
                   → Devuelve JSON
```

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/users/{id}/profile", userProfileHandler.Get)
   ```

**PASO 2: Crear el Handler**

2. Crea: `internal/adapter/http/handler/user_profile_handler.go`
3. Define `UserProfileHandler` con ambos casos de uso:
   ```go
   type UserProfileHandler struct {
       getUserUC     *userUsecase.GetUserUsecase
       listUsersUC   *userUsecase.ListUsersUsecase
   }
   ```

4. Implementa el método `Get`:
   - Extrae y valida el ID
   - Llama a `h.getUserUC.Execute(id)`
   - Llama a `h.listUsersUC.Execute()` y obtiene `len(users)`
   - Construye mensaje: `"Perfil de " + user.Name`
   - Crea respuesta con estructura anónima
   - Devuelve JSON

**PASO 3: Conectar en main.go y router**

5. Crea el handler pasándole ambos casos de uso
6. Regístralo en el router
7. Activa la ruta

**PASO 4: Prueba**

8. `curl http://localhost:8080/users/1/profile`

### ✅ RESULTADO ESPERADO

```json
{
  "user": { "id": 1, "name": "Leanne Graham", ... },
  "total_users": 10,
  "message": "Perfil de Leanne Graham"
}
```

### 💡 LO QUE APRENDISTE

El Handler puede hacer **lógica de presentación** como formatear mensajes y combinar datos sin necesidad de un UseCase adicional.

---

## EJERCICIO 33 - Endpoint de salud detallado (health check)

### 📋 LO QUE NECESITAMOS

Para monitoreo, crea `/health` que verifique:
- Si la API de usuarios responde (intenta obtener usuario 1)
- Si la API de productos responde (intenta obtener producto 1)
- Estado general: "healthy" si ambos funcionan, "degraded" si falla uno, "unhealthy" si fallan ambos

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: El Handler puede usar **UseCases existentes** para verificar salud.

```
HealthHandler → Llama GetUserUsecase(1) → verifica si responde
              → Llama GetProductUsecase(1) → verifica si responde
              → Determina estado según resultados
              → Devuelve JSON
```

**Alternativa válida**: Usar repositorios directamente si necesitas verificar conectividad a bajo nivel.

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/health", healthHandler.Check)
   ```

**PASO 2: Crear el Handler**

2. Crea: `internal/adapter/http/handler/health_handler.go`
3. Define `HealthHandler` que recibe ambos casos de uso:
   ```go
   type HealthHandler struct {
       getUserUC    *userUsecase.GetUserUsecase
       getProductUC *productUsecase.GetProductUsecase
   }
   ```

4. Implementa el método `Check`:
   - Intenta `h.getUserUC.Execute(1)` y verifica si hay error
   - Intenta `h.getProductUC.Execute(1)` y verifica si hay error
   - Calcula el estado:
     ```go
     usersOK := (errUser == nil)
     productsOK := (errProduct == nil)
     
     var status string
     if usersOK && productsOK {
         status = "healthy"
     } else if usersOK || productsOK {
         status = "degraded"
     } else {
         status = "unhealthy"
     }
     ```
   - Construye respuesta con estructura anónima
   - **Siempre devuelve 200 OK** (el health check no debe fallar)

**PASO 3: Conectar en main.go y router**

5. Crea el handler pasándole ambos casos de uso
6. Regístralo y activa la ruta

**PASO 4: Prueba**

7. `curl http://localhost:8080/health`

### ✅ RESULTADO ESPERADO

```json
{
  "status": "healthy",
  "users_api": true,
  "products_api": true,
  "timestamp": "2025-12-15T10:30:00Z"
}
```

### 💡 LO QUE APRENDISTE

Un health check es un caso especial donde el Handler **verifica conectividad** usando UseCases o repositorios existentes. No necesita su propio UseCase dedicado.

---

## EJERCICIO 34 - Endpoint de resumen ejecutivo completo

### 📋 LO QUE NECESITAMOS

Crea `/summary` que devuelva un resumen ejecutivo de toda la API:
- Mensaje de bienvenida
- Versión
- Total de recursos (usuarios + productos)
- Estado de salud (simulado: "healthy")
- Endpoints disponibles
- Última actualización (timestamp)

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: Handler orquesta múltiples casos de uso y agrega datos.

```
SummaryHandler → Llama GetStatusUsecase → extrae versión
               → Llama ListUsersUsecase → cuenta
               → Llama ListProductsUsecase → cuenta
               → Construye lista de endpoints
               → Agrega todo en respuesta
               → Devuelve JSON
```

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta: `// r.Get("/summary", summaryHandler.Get)`

**PASO 2: Crear el Handler (ORQUESTACIÓN COMPLEJA)**

2. Crea: `internal/adapter/http/handler/summary_handler.go`
3. Define `SummaryHandler` con los casos de uso necesarios:
   ```go
   type SummaryHandler struct {
       getStatusUC    *statusUsecase.GetStatusUsecase
       listUsersUC    *userUsecase.ListUsersUsecase
       listProductsUC *productUsecase.ListProductsUsecase
   }
   ```

4. Implementa el método `Get`:
   - Llama a `h.getStatusUC.Execute()` para obtener versión
   - Llama a `h.listUsersUC.Execute()` y cuenta
   - Llama a `h.listProductsUC.Execute()` y cuenta
   - Calcula total: `len(users) + len(products)`
   - Define lista de endpoints (hardcoded o dinámica)
   - Construye respuesta completa con estructura anónima
   - Devuelve JSON

**PASO 3: Conectar en main.go y router**

5. Crea el handler pasándole los tres casos de uso
6. Regístralo y activa la ruta

**PASO 4: Prueba**

7. `curl http://localhost:8080/summary`

### ✅ RESULTADO ESPERADO

```json
{
  "message": "Bienvenido a la API de Ejercicio",
  "version": "1.1.0",
  "total_resources": 30,
  "health_status": "healthy",
  "available_endpoints": ["/status", "/ping", "/users", "/products", ...],
  "timestamp": "2025-12-15T10:30:00Z"
}
```

### 💡 LO QUE APRENDISTE

Endpoints de "resumen" o "dashboard" son perfectos para que el Handler orqueste, ya que solo agregan datos sin lógica de negocio compleja.

---

## EJERCICIO 35 - Endpoint para recomendaciones de productos por usuario

### 📋 LO QUE NECESITAMOS

Aunque no tenemos productos por usuario en las APIs, simularemos esta funcionalidad. Crea `/users/{id}/recommended-products` que devuelva:
- Información del usuario
- 3 productos aleatorios recomendados para ese usuario

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: Handler orquesta y selecciona productos aleatorios.

```
RecommendedProductsHandler → Llama GetUserUsecase(id)
                           → Llama ListProductsUsecase
                           → Selecciona 3 aleatorios (rand)
                           → Construye respuesta
                           → Devuelve JSON
```

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta:
   ```go
   // r.Get("/users/{id}/recommended-products", userProductsHandler.GetRecommended)
   ```

**PASO 2: Crear el Handler**

2. Crea: `internal/adapter/http/handler/user_products_handler.go`
3. Define `UserProductsHandler` con ambos casos de uso:
   ```go
   type UserProductsHandler struct {
       getUserUC       *userUsecase.GetUserUsecase
       listProductsUC  *productUsecase.ListProductsUsecase
   }
   ```

4. Implementa el método `GetRecommended`:
   - Extrae y valida el ID
   - Llama a `h.getUserUC.Execute(id)`
   - Llama a `h.listProductsUC.Execute()`
   - **Selecciona 3 productos aleatorios**:
     ```go
     // Importa "math/rand" y "time"
     rand.Seed(time.Now().UnixNano())
     recommended := make([]*product.Product, 3)
     for i := 0; i < 3 && i < len(products); i++ {
         idx := rand.Intn(len(products))
         recommended[i] = products[idx]
     }
     ```
   - Construye respuesta con estructura anónima
   - Devuelve JSON

**PASO 3: Conectar y Probar**

5. Crea el handler pasándole ambos casos de uso
6. Regístralo y activa la ruta
7. `curl http://localhost:8080/users/1/recommended-products`

### ✅ RESULTADO ESPERADO

El usuario con 3 productos aleatorios cada vez que llames.

### 💡 LO QUE APRENDISTE

Lógica de "recomendación" simple (selección aleatoria) es apropiada para el Handler. Si fuera lógica compleja (ML, análisis de comportamiento), sí justificaría un UseCase dedicado.

---

## EJERCICIO 36 - Endpoint de comparación de productos

### 📋 LO QUE NECESITAMOS

Crea `/products/compare?ids=1,2,3` que permita comparar varios productos lado a lado.

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: Handler orquesta llamadas y calcula métricas simples.

```
CompareProductsHandler → Parse IDs de query string
                       → Llama GetProductUsecase(id) para cada ID
                       → Calcula precio promedio
                       → Identifica más barato/caro
                       → Devuelve JSON
```

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta (debe ir ANTES de `/products/{id}`):
   ```go
   // r.Get("/products/compare", productHandler.Compare)
   ```

**PASO 2: Crear el método en el Handler**

2. En `product_handler.go`
3. Agrega el caso de uso `GetProductUsecase` si no lo tienes ya
4. Crea método `Compare`:
   - Obtiene el parámetro `ids` (string: "1,2,3")
   - Separa por comas: `strings.Split(idsParam, ",")`
   - Convierte cada string a int
   - Valida que haya entre 2 y 4 IDs
   - **Para cada ID, llama** a `h.getProductUC.Execute(id)`
   - **Calcula métricas**:
     ```go
     var total float64
     cheapest := products[0]
     mostExpensive := products[0]
     
     for _, p := range products {
         total += p.Price
         if p.Price < cheapest.Price {
             cheapest = p
         }
         if p.Price > mostExpensive.Price {
             mostExpensive = p
         }
     }
     avgPrice := total / float64(len(products))
     ```
   - Construye respuesta con estructura anónima
   - Devuelve JSON

**PASO 3: Prueba**

5. `curl "http://localhost:8080/products/compare?ids=1,2,3"`

### ✅ RESULTADO ESPERADO

```json
{
  "products": [...],
  "average_price": 150.50,
  "cheapest_id": 2,
  "most_expensive_id": 1
}
```

### 💡 LO QUE APRENDISTE

Cálculos de agregación simples (promedio, mín, máx) son apropiados para el Handler. No necesitas un UseCase separado para esto.

---

## EJERCICIO 37 - Endpoint de dashboard principal

### 📋 LO QUE NECESITAMOS

Crea `/dashboard` que sea el endpoint principal de monitoreo con:
- Estado de salud (simulado: "healthy")
- Total de usuarios y productos
- Top 3 productos más caros
- Timestamp

### 🏛️ ARQUITECTURA CORRECTA

**✅ CORRECTO**: Handler orquesta múltiples UseCases y procesa datos.

```
DashboardHandler → Llama ListUsersUsecase → cuenta
                 → Llama ListProductsUsecase → cuenta + ordena top 3
                 → Simula health status
                 → Construye respuesta agregada
                 → Devuelve JSON
```

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta: `// r.Get("/dashboard", dashboardHandler.Get)`

**PASO 2: Crear el Handler**

2. Crea: `internal/adapter/http/handler/dashboard_handler.go`
3. Define `DashboardHandler` con los casos de uso necesarios
4. Implementa el método `Get`:
   - Llama a `listUsersUC` y cuenta
   - Llama a `listProductsUC`
   - Ordena productos por precio descendente
   - Toma los 3 primeros (top 3 más caros)
   - Construye respuesta con estructura anónima:
     ```go
     response := struct {
         HealthStatus  string              `json:"health_status"`
         TotalUsers    int                 `json:"total_users"`
         TotalProducts int                 `json:"total_products"`
         TopProducts   []*product.Product  `json:"top_3_most_expensive"`
         Timestamp     string              `json:"timestamp"`
     }{...}
     ```
   - Devuelve JSON

**PASO 3: Conectar y Probar**

5. Crea el handler con los casos de uso necesarios
6. Regístralo y activa la ruta
7. `curl http://localhost:8080/dashboard`

### 💡 LO QUE APRENDISTE

Dashboards son el ejemplo perfecto de **orquestación en el Handler**: combinas datos de múltiples fuentes y aplicas transformaciones simples (ordenar, filtrar top N).

---

## EJERCICIO 38 - Endpoint de búsqueda global

### 📋 LO QUE NECESITAMOS

Crea `/search?q=john` que busque en AMBOS recursos (usuarios y productos) y devuelva resultados combinados.

### 🏛️ ARQUITECTURA CORRECTA

**✅ OPCIÓN 1 - Handler orquesta (Recomendado para búsqueda simple)**:
```
SearchHandler → Llama ListUsersUsecase → filtra por término
              → Llama ListProductsUsecase → filtra por término
              → Combina resultados
              → Devuelve JSON
```

**✅ OPCIÓN 2 - UseCase dedicado (Si hay lógica compleja)**:
Si la búsqueda tiene lógica compleja (scoring, relevancia, ponderación), entonces SÍ justifica un `SearchUsecase` porque es una **regla de negocio real**.

Para este ejercicio: **usa Opción 1** (Handler simple).

### 🎯 INSTRUCCIONES

**PASO 1: Registrar la ruta**

1. En `router.go`, comenta: `// r.Get("/search", searchHandler.Search)`

**PASO 2: Crear el Handler**

2. Crea: `internal/adapter/http/handler/search_handler.go`
3. Define `SearchHandler` con ambos casos de uso
4. Implementa el método `Search`:
   - Obtiene parámetro `q` de la query
   - Valida que no esté vacío
   - Llama a `h.listUsersUC.Execute()`
   - **Filtra usuarios** que contengan el término en nombre, email o username:
     ```go
     term := strings.ToLower(query)
     matchedUsers := []*user.User{}
     for _, u := range users {
         if strings.Contains(strings.ToLower(u.Name), term) ||
            strings.Contains(strings.ToLower(u.Email), term) ||
            strings.Contains(strings.ToLower(u.Username), term) {
             matchedUsers = append(matchedUsers, u)
         }
     }
     ```
   - Llama a `h.listProductsUC.Execute()`
   - **Filtra productos** que contengan el término en título o descripción
   - Calcula total: `len(matchedUsers) + len(matchedProducts)`
   - Construye respuesta con estructura anónima
   - Devuelve JSON

**PASO 3: Conectar y Probar**

5. Crea el handler con ambos casos de uso
6. Regístralo y activa la ruta
7. `curl "http://localhost:8080/search?q=shirt"`

### ✅ RESULTADO ESPERADO

```json
{
  "search_term": "shirt",
  "matched_users": [],
  "matched_products": [productos con "shirt"],
  "total_results": 3
}
```

### 💡 LO QUE APRENDISTE

**Búsqueda simple** (contiene texto) → Handler lo hace.  
**Búsqueda compleja** (scoring, ML, relevancia) → Sí justifica un UseCase dedicado con lógica de negocio real.

---

## 🎓 FELICIDADES - HAS COMPLETADO EL TALLER

Has completado **38 ejercicios progresivos** de desarrollo de APIs REST. Ahora dominas:

### 📝 NIVEL BÁSICO (Ejercicios 1-8)
✅ Cambiar textos y valores en endpoints existentes
✅ Agregar validaciones de datos (IDs, rangos, tipos)
✅ Agregar nuevos campos a estructuras existentes
✅ Mejorar mensajes de error específicos por código HTTP

### 🏗️ NIVEL INTERMEDIO (Ejercicios 9-30)
✅ Crear endpoints completos desde cero siguiendo el flujo correcto
✅ Crear módulos completos (productos) con todas las capas
✅ Implementar búsquedas con query parameters
✅ Crear endpoints especializados (transformaciones parciales)
✅ Filtrar datos por múltiples criterios (categoría, precio, texto, rangos)
✅ Reutilizar casos de uso existentes eficientemente
✅ Transformar y agregar datos en memoria (conteos, resúmenes)
✅ Implementar paginación básica con metadatos
✅ Ordenar colecciones por diferentes criterios dinámicamente
✅ Extraer valores únicos de colecciones
✅ Validar datos con expresiones regulares (regex)
✅ Combinar múltiples filtros opcionales
✅ Manejar errores con graceful degradation (valores por defecto)
✅ Aplicar transformaciones complejas de datos (formateo, cálculos)
✅ Trabajar con query parameters opcionales
✅ Manipular strings (truncar, formatear, convertir caso)

### 🔥 NIVEL AVANZADO (Ejercicios 31-38)
✅ **Orquestar múltiples UseCases desde el Handler** (arquitectura correcta)
✅ Crear endpoints de estadísticas con agregación en el Handler
✅ Implementar health checks verificando dependencias externas
✅ Construir dashboards combinando datos de múltiples fuentes
✅ Aplicar transformaciones y cálculos simples en el Handler
✅ Filtrar y procesar datos de diferentes dominios
✅ Crear recomendaciones y comparaciones con lógica simple
✅ Implementar búsqueda global multi-recurso
✅ Gestionar respuestas complejas sin acoplar UseCases
✅ Diseñar endpoints para monitoreo y observabilidad
✅ **Entender cuándo SÍ y cuándo NO crear un UseCase dedicado**

### 🎯 ARQUITECTURA Y MEJORES PRÁCTICAS
✅ **Flujo correcto**: Ruta → Handler → Caso de Uso → Repositorio → Dominio
✅ **Separación de responsabilidades**: Cada capa con su función específica
✅ **Reutilización de código**: Compartir casos de uso entre handlers
✅ **Validación en capas**: Formato en handler, negocio en caso de uso
✅ **Manejo de errores**: Mensajes claros y específicos
✅ **DRY**: No repetir código, extraer funcionalidad común
✅ **Testing**: Verificar cada capa independientemente

## 📊 PROGRESIÓN DE APRENDIZAJE

### Mapa Completo del Taller

```
NIVEL BÁSICO (1-8) - ⭐
├─ Ejercicio 1: Cambiar mensaje del ping
├─ Ejercicio 2: Cambiar versión del status
├─ Ejercicio 3: Validar IDs negativos
├─ Ejercicio 4: Limitar rango de IDs
├─ Ejercicio 5: Rechazar IDs no numéricos
├─ Ejercicio 6: Agregar campo environment
├─ Ejercicio 7: Agregar timestamp
└─ Ejercicio 8: Mejorar mensajes de error 404

NIVEL INTERMEDIO INICIAL (9-14) - ⭐⭐
├─ Ejercicio 9: Listar todos los usuarios ← PRIMER ENDPOINT COMPLETO
├─ Ejercicio 10: Endpoint de bienvenida
├─ Ejercicio 11: Contador de peticiones
├─ Ejercicio 12: Combinar usuario + status
├─ Ejercicio 13: Utilidad de validación
└─ Ejercicio 14: Módulo completo de productos ← MÓDULO DESDE CERO

NIVEL INTERMEDIO - FILTROS Y BÚSQUEDAS (15-22) - ⭐⭐
├─ Ejercicio 15: Buscar usuarios por email ← MEJORADO (con contexto de negocio + arquitectura + validaciones + malas prácticas)
├─ Ejercicio 16: Endpoint /users/{id}/email (reutilización)
├─ Ejercicio 17: Filtrar productos por categoría
├─ Ejercicio 18: Productos con precio mayor a X
├─ Ejercicio 19: Filtrar usuarios por dominio de email ← SIMPLIFICADO
├─ Ejercicio 20: Listar solo títulos de productos
├─ Ejercicio 21: Contar productos por categoría
└─ Ejercicio 22: Buscar productos por título

NIVEL INTERMEDIO - ORDENAMIENTO Y AGREGACIÓN (23-30) - ⭐⭐ a ⭐⭐⭐
├─ Ejercicio 23: Ordenar productos por precio ← SIMPLIFICADO (solo ascendente)
├─ Ejercicio 24: Ordenar productos por precio descendente
├─ Ejercicio 25: Combinar filtros (categoría + precio mínimo) ← SIMPLIFICADO
├─ Ejercicio 26: Producto más barato ← SIMPLIFICADO
├─ Ejercicio 27: Filtrar por rango de precio ← SIMPLIFICADO
├─ Ejercicio 28: Obtener ciudad de usuario (datos anidados) ← SIMPLIFICADO
├─ Ejercicio 29: Filtrar por rango de precio ampliado
└─ Ejercicio 30: Listar categorías únicas

NIVEL AVANZADO (31-38) - ⭐⭐⭐
├─ Ejercicio 31: Estadísticas de la API ← HANDLER ORQUESTA MÚLTIPLES CASOS DE USO
├─ Ejercicio 32: Perfil completo de usuario ← HANDLER ORQUESTA Y FORMATEA
├─ Ejercicio 33: Health check detallado ← HANDLER VERIFICA SALUD DE DEPENDENCIAS
├─ Ejercicio 34: Resumen ejecutivo completo ← HANDLER AGREGA DATOS DE MÚLTIPLES FUENTES
├─ Ejercicio 35: Recomendaciones de productos ← HANDLER SELECCIONA ALEATORIAMENTE
├─ Ejercicio 36: Comparación de productos ← HANDLER CALCULA MÉTRICAS
├─ Ejercicio 37: Dashboard principal ← HANDLER ORQUESTA Y PROCESA
└─ Ejercicio 38: Búsqueda global multi-recurso ← HANDLER FILTRA EN MÚLTIPLES DOMINIOS

TOTAL: 38 ejercicios progresivos
```

### Resumen por Niveles

| Nivel | Ejercicios | Habilidades | Complejidad |
|-------|-----------|-------------|-------------|
| **Básico** | 1-8 | Modificar código existente, validaciones simples | ⭐ |
| **Intermedio** | 9-30 | Crear endpoints, filtros, transformaciones | ⭐⭐ |
| **Avanzado** | 31-38 | Múltiples dominios, agregaciones complejas | ⭐⭐⭐ |

## 🚀 SIGUIENTES DESAFÍOS

### Nivel Implementación Real
1. **Agrega paginación** a los endpoints de listas (limit, offset, total)
2. **Combina filtros múltiples** (`/products?category=electronics&minPrice=100&maxPrice=500`)
3. **Agrega ordenamiento dinámico** (`/products?sortBy=price&order=desc`)
4. **Implementa caché** (Redis o in-memory) para reducir llamadas externas
5. **Agrega rate limiting** para proteger endpoints de abuso

### Nivel Profesional
6. **Autenticación JWT** para proteger endpoints privados
7. **Base de datos real** (PostgreSQL/MySQL) en lugar de APIs externas
8. **Endpoints de escritura** (POST/PUT/DELETE) con validación completa
9. **Validación con schemas** (go-validator, JSON Schema)
10. **Documentación automática** con Swagger/OpenAPI
11. **Versionado de API** (v1, v2) con compatibilidad hacia atrás
12. **CORS y headers de seguridad** configurados correctamente

### Nivel Experto
13. **Testing completo**: unitarios, integración y E2E
14. **Observabilidad**: logs estructurados, métricas (Prometheus), tracing (Jaeger)
15. **Circuit breakers** y timeouts para APIs externas
16. **Retry policies** con backoff exponencial
17. **Graceful shutdown** y manejo de señales
18. **Soporte multiidioma** (i18n) en mensajes de error
19. **Webhooks** para notificaciones asíncronas
20. **GraphQL** como alternativa a REST
21. **Containerización** con Docker y docker-compose
22. **CI/CD** con GitHub Actions o GitLab CI
23. **Despliegue en Cloud** (AWS, GCP, Azure)

## 💡 CONCEPTOS CLAVE DOMINADOS

### 🔄 El Flujo Cronológico (LO MÁS IMPORTANTE)

```
SIEMPRE SEGUIR ESTE ORDEN:

1️⃣ RUTA (router.go)
   "Defino QUÉ endpoint quiero exponer"
   
2️⃣ HANDLER (handler/)
   "Recibo la petición HTTP"
   "Extraigo parámetros de URL/query"
   "Valido formato (¿es un número? ¿está vacío?)"
   
3️⃣ CASO DE USO (usecase/)
   "Aplico lógica de negocio"
   "Valido reglas (¿ID entre 1-10? ¿email válido?)"
   "Llamo al repositorio"
   
4️⃣ REPOSITORIO (adapter/repository/)
   "Obtengo datos de API externa o BD"
   "Manejo errores de conexión"
   
5️⃣ DOMINIO (domain/)
   "Define estructuras y contratos"
   "Ya suele estar creado"
```

### 🎯 Arquitectura y Patrones

1. **Arquitectura en capas**: Separación clara entre presentación, lógica y datos
2. **Dependency Injection**: Pasar dependencias en constructores
3. **Interfaces**: Contratos entre capas para flexibilidad
4. **Error handling**: Propagación y transformación de errores
5. **HTTP Status Codes**: Uso correcto según el contexto
6. **REST Best Practices**: Nombres de recursos, verbos HTTP, estructuras de respuesta
7. **Code organization**: Estructura de carpetas escalable
8. **Reusabilidad**: Evitar duplicación, extraer funcionalidad común
9. **Testability**: Código fácil de testear por su diseño
10. **Mantenibilidad**: Código claro, consistente y bien documentado

### 📝 Recordatorios Importantes

✅ **SIEMPRE** empieza por la ruta  
✅ **NUNCA** empieces por el repositorio o dominio  
✅ **PRIMERO** piensa en QUÉ necesitas, luego en CÓMO lo obtienes  
✅ **REUTILIZA** casos de uso existentes cuando sea posible  
✅ **VALIDA** en dos capas: formato en handler, negocio en caso de uso  
✅ **PRUEBA** después de cada paso importante  
✅ **LEE** el ejercicio completo antes de empezar

## 🎖️ HABILIDADES ADQUIRIDAS

Ahora puedes:
- ✅ Diseñar y crear APIs REST completas desde cero
- ✅ Implementar arquitectura hexagonal / clean architecture
- ✅ Manejar múltiples fuentes de datos
- ✅ Crear endpoints para diferentes casos de uso
- ✅ Validar datos en múltiples capas
- ✅ Implementar patrones de diseño comunes
- ✅ Estructurar proyectos de forma escalable
- ✅ Debuggear y resolver problemas eficientemente
- ✅ Leer y entender código de otros desarrolladores
- ✅ Contribuir a proyectos Go de forma profesional

## 🌟 PRÓXIMOS PASOS RECOMENDADOS

1. **Practica regularmente**: Crea tu propio proyecto desde cero
2. **Lee código**: Explora proyectos open source en Go
3. **Contribuye**: Participa en proyectos de código abierto
4. **Aprende patrones**: Design patterns, architectural patterns
5. **Profundiza en Go**: Goroutines, channels, context, testing avanzado
6. **Explora frameworks**: Echo, Gin, Fiber (aunque ya dominas lo fundamental)
7. **Aprende sobre infraestructura**: Docker, Kubernetes, cloud services

¡Felicidades por completar el taller! Ahora tienes una base sólida para desarrollar APIs profesionales en Go. 🎉

---

## ✅ CHECKLIST RÁPIDA PARA CADA EJERCICIO

Usa esta lista cada vez que hagas un ejercicio:

### Antes de Empezar
- [ ] Leí el ejercicio completo
- [ ] Entiendo QUÉ necesito crear
- [ ] Sé en qué nivel estoy (básico/intermedio/avanzado)
- [ ] El servidor está corriendo en otra terminal

### Durante el Desarrollo (sigue este orden)
- [ ] **PASO 1**: Comenté la ruta en `router.go`
- [ ] **PASO 2**: Creé/modifiqué el Handler
- [ ] **PASO 3**: Creé el Caso de Uso (si es necesario)
- [ ] **PASO 4**: Definí el Dominio (si es necesario)
- [ ] **PASO 5**: Implementé el Repositorio (si es necesario)
- [ ] **PASO 6**: Conecté todo en el Router
- [ ] Revisé que no haya errores de compilación
- [ ] Los imports están correctos

### Después de Implementar
- [ ] Reinicié el servidor
- [ ] Probé el endpoint con curl
- [ ] Verifiqué el resultado esperado
- [ ] Probé casos de error (IDs inválidos, parámetros faltantes)
- [ ] El servidor sigue funcionando sin errores

### Si Algo No Funciona
- [ ] Leí el mensaje de error completo
- [ ] Revisé los nombres (mayúsculas/minúsculas)
- [ ] Verifiqué los imports
- [ ] Comparé con ejemplos de ejercicios anteriores
- [ ] Revisé que seguí el orden: Ruta → Handler → Caso de Uso → Repositorio

---

**¡Sigue construyendo y aprendiendo!** 🚀


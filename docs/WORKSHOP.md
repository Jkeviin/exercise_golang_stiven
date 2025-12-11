# 🎓 TALLER PRÁCTICO - Desarrollo de APIs REST

Este taller te guiará para agregar nuevas funcionalidades a la API paso a paso.

**Importante**: Cada ejercicio explica QUÉ necesitas agregar, no CÓMO programarlo.

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

## EJERCICIO 1 - Cambiar el mensaje de respuesta de ping

### 📋 QUÉ QUEREMOS

El endpoint `/ping` actualmente responde con `{"message":"pong"}`. Queremos que responda con `{"message":"¡Servidor activo!"}`.

### 🎯 PASOS DETALLADOS

1. **Encuentra el archivo donde se define qué responde el endpoint ping**
   - Busca en la carpeta `internal/usecase/ping/`
   - Abre el archivo que tiene el nombre relacionado con "ping"
   - Dentro verás una línea que dice `Message: "pong"`

2. **Cambia el mensaje**
   - Reemplaza el texto `"pong"` por `"¡Servidor activo!"`
   - Guarda el archivo

3. **Prueba el cambio**
   - Detén el servidor (presiona Ctrl+C en la terminal donde está corriendo)
   - Vuelve a iniciar el servidor: `go run cmd/app/main.go`
   - En otra terminal ejecuta: `curl http://localhost:8080/ping`
   - Deberías ver: `{"message":"¡Servidor activo!"}`

### ✅ VERIFICACIÓN

- ✅ El endpoint `/ping` ahora responde con el nuevo mensaje
- ✅ El endpoint `/status` sigue funcionando igual
- ✅ El endpoint `/users/1` sigue funcionando igual

### 💡 QUÉ APRENDISTE

Has modificado la respuesta de un endpoint cambiando el mensaje en el caso de uso.

---

## EJERCICIO 2 - Cambiar la versión de la aplicación

### 📋 QUÉ QUEREMOS

El endpoint `/status` muestra `"version":"1.0.0"`. Queremos cambiarla a `"version":"1.1.0"` porque hicimos mejoras.

### 🎯 PASOS DETALLADOS

1. **Encuentra el archivo donde se define la versión**
   - Busca en la carpeta `internal/usecase/status/`
   - Abre el archivo relacionado con "status"
   - Busca la línea que dice `Version: "1.0.0"`

2. **Cambia la versión**
   - Reemplaza `"1.0.0"` por `"1.1.0"`
   - Guarda el archivo

3. **Prueba el cambio**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - La respuesta debe mostrar `"version":"1.1.0"`

### ✅ VERIFICACIÓN

- ✅ El campo `version` ahora muestra `1.1.0`
- ✅ Los demás campos siguen igual (`message`, `uptime`)

### 💡 QUÉ APRENDISTE

Has modificado un dato que devuelve el endpoint cambiando el valor en el caso de uso.

---

## EJERCICIO 3 - Agregar un nuevo campo al status

### 📋 QUÉ QUEREMOS

El endpoint `/status` actualmente devuelve `message`, `version` y `uptime`. Queremos agregar un campo nuevo llamado `environment` que diga `"development"`.

### 🎯 PASOS DETALLADOS

1. **Agrega el campo a la estructura de datos**
   - Busca la carpeta `internal/domain/status/`
   - Abre el archivo de la entidad Status
   - Verás campos como `Message string`, `Version string`, `Uptime int64`
   - Agrega un nuevo campo después de los existentes:
     - Nombre del campo: `Environment`
     - Tipo: `string`
     - Tag JSON: `"environment"`
   - El formato debe ser igual a los otros campos que ya existen

2. **Haz que el caso de uso devuelva el nuevo campo**
   - Ve a `internal/usecase/status/`
   - Abre el archivo del caso de uso
   - Busca donde se crea el objeto Status (donde dice `Message: "..."`, `Version: "..."`, etc.)
   - Agrega una nueva línea para el campo `Environment` con el valor `"development"`

3. **Prueba el cambio**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - Ahora deberías ver un campo adicional: `"environment":"development"`

### ✅ VERIFICACIÓN

La respuesta debe verse así:
```json
{
  "message": "...",
  "version": "1.1.0",
  "uptime": 5,
  "environment": "development"
}
```

### 💡 QUÉ APRENDISTE

Para agregar un nuevo campo a una respuesta:
1. Lo agregas a la estructura de datos (entidad)
2. Lo llenas con un valor en el caso de uso

---

## EJERCICIO 4 - Agregar timestamp al status

### 📋 QUÉ QUEREMOS

Queremos que el endpoint `/status` también devuelva la fecha y hora actual del servidor en un campo llamado `timestamp`.

### 🎯 PASOS DETALLADOS

1. **Agrega el campo timestamp a la entidad**
   - Ve a `internal/domain/status/`
   - En la estructura Status, agrega un nuevo campo:
     - Nombre: `Timestamp`
     - Tipo: `string`
     - Tag JSON: `"timestamp"`

2. **Haz que el caso de uso genere el timestamp**
   - Ve a `internal/usecase/status/`
   - Busca donde se crea el objeto Status
   - Agrega el campo `Timestamp` con el valor de la fecha/hora actual
   - Usa el formato: `time.Now().Format(time.RFC3339)`
   - **Nota**: Si ves error de "undefined: time", agrega `"time"` en los imports del archivo

3. **Prueba el cambio**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/status`
   - Deberías ver un campo `timestamp` con formato: `"2025-12-11T14:30:45Z"`

### ✅ VERIFICACIÓN

La respuesta debe incluir:
```json
{
  "message": "...",
  "version": "1.1.0",
  "uptime": 3,
  "environment": "development",
  "timestamp": "2025-12-11T14:30:45Z"
}
```

### 💡 QUÉ APRENDISTE

Puedes usar funciones del sistema (como obtener la hora actual) dentro de los casos de uso.

---

## EJERCICIO 5 - Validar que el ID de usuario no sea mayor a 10

### 📋 QUÉ QUEREMOS

Actualmente puedes llamar `/users/999` y el servidor intenta buscar ese usuario. Queremos que si el ID es mayor a 10, el servidor responda con un error diciendo "El ID debe estar entre 1 y 10".

### 🎯 PASOS DETALLADOS

1. **Encuentra dónde está la validación actual del ID**
   - Ve a `internal/usecase/user/`
   - Abre el archivo del caso de uso de obtener usuario
   - Busca la parte que valida `if id <= 0`
   - Esa línea verifica que el ID sea positivo

2. **Agrega la nueva validación**
   - Después de la validación existente, agrega una nueva condición:
   - Si el ID es mayor que 10, devuelve un error
   - El mensaje debe ser: `"el ID debe estar entre 1 y 10"`
   - Usa el mismo formato que la validación existente

3. **Prueba la validación**
   - Reinicia el servidor
   - Prueba con ID válido: `curl http://localhost:8080/users/5`
     - Debe funcionar normalmente
   - Prueba con ID muy grande: `curl http://localhost:8080/users/999`
     - Debe dar error con el mensaje que configuraste
   - Prueba con ID 0: `curl http://localhost:8080/users/0`
     - Debe dar el error original "el ID debe ser mayor que 0"

### ✅ VERIFICACIÓN

- ✅ `/users/1` a `/users/10` funcionan
- ✅ `/users/11` o mayor da error "el ID debe estar entre 1 y 10"
- ✅ `/users/0` o negativo da error "el ID debe ser mayor que 0"

### 💡 QUÉ APRENDISTE

Las validaciones de negocio (como rangos permitidos) se ponen en el caso de uso.

---

## EJERCICIO 6 - Crear endpoint para listar todos los usuarios

### 📋 QUÉ QUEREMOS

Tenemos el endpoint `/users/{id}` que devuelve un usuario. Queremos crear un nuevo endpoint `/users` (sin ID) que devuelva la lista de todos los usuarios.

### 🎯 PASOS DETALLADOS

#### PARTE A: Actualizar el repositorio

1. **Agrega el método a la interfaz del repositorio**
   - Ve a `internal/domain/user/`
   - Abre el archivo que define la interfaz `Repository`
   - Actualmente tiene un método: `FindByID(id int) (*User, error)`
   - Agrega un nuevo método debajo:
     - Nombre: `FindAll`
     - No recibe parámetros
     - Devuelve: `([]*User, error)` (una lista de usuarios y posible error)

2. **Implementa el nuevo método en el repositorio**
   - Ve a `internal/adapter/repository/`
   - Abre el archivo que implementa el repositorio de usuarios
   - Busca el método `FindByID` para ver cómo está hecho
   - Crea un nuevo método `FindAll` usando el mismo patrón:
     - URL: `{baseURL}/users` (sin el ID al final)
     - Decodifica en un slice de usuarios: `var users []*user.User`
     - Devuelve la lista

#### PARTE B: Crear el caso de uso

3. **Crea un nuevo archivo para el caso de uso**
   - En la carpeta `internal/usecase/user/`
   - Crea un archivo nuevo llamado `list_users.go`
   - Copia la estructura del archivo `get_user.go` pero:
     - Cambia el nombre a `ListUsersUsecase`
     - El método se llama `Execute()` (sin parámetros)
     - Devuelve `([]*user.User, error)`
     - Solo llama a `userRepo.FindAll()` sin validaciones

#### PARTE C: Crear el handler

4. **Agrega el nuevo handler**
   - Ve a `internal/adapter/http/handler/`
   - Abre el archivo `user_handler.go`
   - Ya existe un método `GetByID`, vamos a agregar otro
   - Primero, modifica la estructura `UserHandler` para agregar el nuevo caso de uso:
     - Agrega un campo `listUsersUC *userUsecase.ListUsersUsecase`
   - Modifica el constructor para recibir ambos casos de uso
   - Crea un nuevo método `List(w http.ResponseWriter, r *http.Request)`:
     - Llama a `h.listUsersUC.Execute()`
     - Devuelve la lista en JSON (igual que hace `GetByID`)

#### PARTE D: Registrar la ruta

5. **Conecta el nuevo endpoint en el router**
   - Ve a `internal/infrastructure/http/`
   - Abre el archivo `router.go`
   - Busca donde se crea el `userHandler` (donde dice `NewUserHandler`)
   - Primero, crea el nuevo caso de uso antes de crear el handler:
     - `listUsersUC := userUsecase.NewListUsersUsecase(userRepo)`
   - Modifica la línea donde se crea el handler para pasar ambos casos de uso
   - Después de la línea `r.Get("/users/{id}", userHandler.GetByID)`
   - Agrega: `r.Get("/users", userHandler.List)`

6. **Prueba el nuevo endpoint**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/users`
   - Deberías ver una lista de 10 usuarios

### ✅ VERIFICACIÓN

- ✅ `/users` devuelve una lista de usuarios (array JSON)
- ✅ `/users/1` sigue funcionando (un solo usuario)

### 💡 QUÉ APRENDISTE

Para crear un endpoint completo:
1. Agregas el método a la interfaz del repositorio
2. Lo implementas en el repositorio concreto
3. Creas un caso de uso que lo usa
4. Creas un handler que llama al caso de uso
5. Registras la ruta en el router

---

## EJERCICIO 7 - Agregar endpoint de bienvenida

### 📋 QUÉ QUEREMOS

Crear un nuevo endpoint `/` (ruta raíz) que devuelva un mensaje de bienvenida cuando alguien entre a la API.

Respuesta esperada:
```json
{
  "message": "Bienvenido a la API de Ejercicio",
  "version": "1.1.0",
  "endpoints": [
    "/status",
    "/ping",
    "/users",
    "/users/{id}"
  ]
}
```

### 🎯 PASOS DETALLADOS

1. **Crea la entidad para la bienvenida**
   - Crea una carpeta nueva en `internal/domain/` llamada `welcome`
   - Dentro crea un archivo `welcome.go`
   - Define una estructura con los campos:
     - `Message` (string)
     - `Version` (string)
     - `Endpoints` (slice de strings: `[]string`)

2. **Crea el caso de uso**
   - Crea una carpeta `internal/usecase/welcome/`
   - Crea el archivo `get_welcome.go`
   - Crea el caso de uso `GetWelcomeUsecase`
   - El método `Execute()` debe devolver la estructura Welcome con:
     - Message: "Bienvenido a la API de Ejercicio"
     - Version: "1.1.0"
     - Endpoints: la lista de endpoints (como array)

3. **Crea el handler**
   - En `internal/adapter/http/handler/`
   - Crea el archivo `welcome_handler.go`
   - Sigue el mismo patrón que `ping_handler.go`
   - El método debe llamar al caso de uso y devolver JSON

4. **Registra la ruta**
   - En `router.go`, al inicio (después de crear el router)
   - Crea el caso de uso, el handler y registra la ruta `/`

5. **Prueba**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/`
   - Deberías ver el mensaje de bienvenida completo

### ✅ VERIFICACIÓN

- ✅ `curl http://localhost:8080/` muestra la bienvenida
- ✅ Muestra la versión correcta
- ✅ Muestra la lista de endpoints

---

## EJERCICIO 8 - Mejorar mensajes de error del repositorio

### 📋 QUÉ QUEREMOS

Cuando el endpoint de usuarios falla (ejemplo: usuario no existe), el mensaje de error no es muy claro. Queremos mejorar los mensajes según el código de error que devuelve la API externa.

### 🎯 PASOS DETALLADOS

1. **Mejora el manejo de errores en el repositorio**
   - Ve a `internal/adapter/repository/`
   - Abre el archivo del repositorio de usuarios
   - Busca la parte que valida `resp.StatusCode != http.StatusOK`
   - Reemplázala por múltiples validaciones:
     - Si `resp.StatusCode == 404`: error "usuario no encontrado"
     - Si `resp.StatusCode >= 500`: error "el servidor externo no está disponible"
     - Si otro código: error "error inesperado del servidor: código {código}"

2. **Prueba los diferentes casos de error**
   - ID que no existe: `curl http://localhost:8080/users/999`
     - Debe decir "usuario no encontrado"
   - Si la API externa falla (simula apagando tu internet un momento)
     - Debe decir error de conexión

### ✅ VERIFICACIÓN

- ✅ Errores 404 muestran "usuario no encontrado"
- ✅ Errores 500+ muestran "servidor externo no disponible"
- ✅ IDs válidos siguen funcionando

---

## EJERCICIO 9 - Crear endpoint combinado usuario + estado

### 📋 QUÉ QUEREMOS

Crear un endpoint `/user-info/{id}` que devuelva en una sola respuesta:
- La información del usuario
- El estado actual del servidor

Esto es útil cuando el cliente necesita ambos datos y no quiere hacer dos peticiones separadas.

Respuesta esperada:
```json
{
  "user": {
    "id": 1,
    "name": "Leanne Graham",
    "email": "...",
    "username": "..."
  },
  "server_status": {
    "message": "...",
    "version": "1.1.0",
    "uptime": 42,
    "environment": "development",
    "timestamp": "..."
  }
}
```

### 🎯 PASOS DETALLADOS

1. **Crea la nueva entidad combinada**
   - Crea carpeta `internal/domain/userinfo/`
   - Crea archivo `user_info.go`
   - Define estructura `UserInfo` con dos campos:
     - `User` de tipo `*user.User`
     - `ServerStatus` de tipo `*status.Status`
   - Tags JSON: `"user"` y `"server_status"`

2. **Crea el caso de uso combinado**
   - Crea carpeta `internal/usecase/userinfo/`
   - Crea archivo `get_user_info.go`
   - Este caso de uso necesita DOS dependencias:
     - El caso de uso de GetUser
     - El caso de uso de GetStatus
   - El método `Execute(id int)` debe:
     - Llamar a GetUser con el ID
     - Llamar a GetStatus
     - Combinar ambos en UserInfo
     - Devolver el resultado

3. **Crea el handler**
   - En `internal/adapter/http/handler/`
   - Crea `user_info_handler.go`
   - Extrae el ID de la URL (como hace `user_handler.go`)
   - Llama al caso de uso con el ID
   - Devuelve el JSON combinado

4. **Registra la ruta**
   - En `router.go`, después de las otras rutas de users
   - Crea el caso de uso combinado pasándole los dos casos de uso que necesita
   - Crea el handler
   - Registra la ruta `/user-info/{id}`

5. **Prueba**
   - Reinicia el servidor
   - Ejecuta: `curl http://localhost:8080/user-info/1`
   - Deberías ver ambos datos combinados

### ✅ VERIFICACIÓN

- ✅ El endpoint devuelve tanto el usuario como el status
- ✅ Si el ID es inválido, solo muestra error del usuario
- ✅ Los endpoints originales siguen funcionando

### 💡 QUÉ APRENDISTE

Puedes crear casos de uso que usan otros casos de uso para combinar funcionalidades.

---

## EJERCICIO 10 - Agregar test para validación de ID

### 📋 QUÉ QUEREMOS

Crear una prueba automática que verifique que la validación de IDs funciona correctamente. Esto asegura que si alguien modifica el código en el futuro, la validación siga funcionando.

### 🎯 PASOS DETALLADOS

1. **Abre el archivo de tests existente**
   - Ve a `test/usecase/user/`
   - Abre el archivo de tests de usuario
   - Verás que ya hay tests como `TestGetUserUsecase_Execute`

2. **Agrega un nuevo test para validar el límite superior**
   - Crea una nueva función de test (copia el formato de los existentes)
   - Nombre: `TestGetUserUsecase_Execute_IDTooHigh`
   - Dentro del test:
     - Crea el mock del repositorio (igual que los otros tests)
     - Crea el caso de uso
     - Llama a `Execute(99)` (un ID mayor a 10)
     - Verifica que SÍ devuelve error
     - Verifica que el mensaje del error contiene "entre 1 y 10"

3. **Ejecuta los tests**
   - En la terminal: `go test ./test/... -v`
   - Todos los tests deben pasar
   - Deberías ver tu nuevo test en la lista

### ✅ VERIFICACIÓN

- ✅ El comando de tests muestra todos los tests pasando
- ✅ Aparece tu nuevo test en la lista
- ✅ Si cambias la validación del caso de uso, el test lo detecta

---

## EJERCICIO 11 - Agregar contador de peticiones al status

### 📋 QUÉ QUEREMOS

Queremos que el endpoint `/status` también muestre cuántas veces ha sido llamado desde que arrancó el servidor.

Ejemplo: Si llamas `/status` 3 veces, la tercera vez debe mostrar `"request_count": 3`.

### 🎯 PASOS DETALLADOS

1. **Agrega el campo a la entidad**
   - En `internal/domain/status/`
   - Agrega campo `RequestCount` (tipo `int`) con tag JSON `"request_count"`

2. **Agrega un contador al caso de uso**
   - En `internal/usecase/status/`
   - La estructura `GetStatusUsecase` ya tiene un campo `startTime`
   - Agrega otro campo llamado `requestCount` (tipo `int`)
   - En el método `Execute()`:
     - Incrementa el contador: `uc.requestCount++`
     - Incluye el contador en la respuesta: `RequestCount: uc.requestCount`

3. **Prueba**
   - Reinicia el servidor
   - Llama varias veces: `curl http://localhost:8080/status`
   - El número debe aumentar en cada llamada

### ✅ VERIFICACIÓN

- Primera llamada: `"request_count": 1`
- Segunda llamada: `"request_count": 2`
- Tercera llamada: `"request_count": 3`

### 💡 QUÉ APRENDISTE

Los casos de uso pueden mantener estado interno (variables que cambian con cada llamada).

---

## EJERCICIO 12 - Crear módulo completo de productos

### 📋 QUÉ QUEREMOS

Crear un módulo completamente nuevo para manejar productos. La API de productos está en `https://fakestoreapi.com`.

Necesitamos:
- Endpoint `/products/{id}` para obtener un producto por ID
- Endpoint `/products` para listar todos los productos

Un producto tiene:
- ID (número)
- Title (texto)
- Price (número decimal)
- Description (texto)
- Category (texto)

### 🎯 PASOS DETALLADOS

#### PARTE 1: Crear el dominio

1. **Crea la entidad Product**
   - Carpeta: `internal/domain/product/`
   - Archivo: `product.go`
   - Estructura con los campos mencionados arriba

2. **Crea la interfaz del repositorio**
   - Archivo: `internal/domain/product/repository.go`
   - Métodos:
     - `FindByID(id int) (*Product, error)`
     - `FindAll() ([]*Product, error)`

#### PARTE 2: Implementar el repositorio

3. **Crea el repositorio que consulta la API externa**
   - Carpeta: `internal/adapter/repository/`
   - Archivo: `product_api_repository.go`
   - URL base: `https://fakestoreapi.com`
   - Implementa ambos métodos siguiendo el patrón de `user_api_repository.go`

#### PARTE 3: Crear los casos de uso

4. **Caso de uso para obtener un producto**
   - Carpeta: `internal/usecase/product/`
   - Archivo: `get_product.go`
   - Valida que el ID esté entre 1 y 20

5. **Caso de uso para listar productos**
   - Archivo: `list_products.go`
   - No necesita validaciones, solo llama al repositorio

#### PARTE 4: Crear los handlers

6. **Crea el handler de productos**
   - Carpeta: `internal/adapter/http/handler/`
   - Archivo: `product_handler.go`
   - Dos métodos: `GetByID` y `List`

#### PARTE 5: Registrar las rutas

7. **Conecta todo en el router**
   - En `router.go`, después de las rutas de users
   - Crea el repositorio, los casos de uso, el handler
   - Registra:
     - `GET /products`
     - `GET /products/{id}`

#### PARTE 6: Crear tests

8. **Crea tests para los casos de uso**
   - Carpeta: `test/usecase/product/`
   - Archivos:
     - `get_product_test.go`
     - `list_products_test.go`
   - Usa mocks como en los tests de user

### ✅ VERIFICACIÓN

- ✅ `curl http://localhost:8080/products` muestra lista de productos
- ✅ `curl http://localhost:8080/products/1` muestra un producto
- ✅ `curl http://localhost:8080/products/999` da error de validación
- ✅ Los tests pasan: `go test ./test/... -v`

### 💡 QUÉ APRENDISTE

Has creado un módulo completo desde cero siguiendo el mismo patrón de la arquitectura.

---

## 🎓 FELICIDADES

Has completado el taller. Ahora sabes:

✅ Modificar respuestas de endpoints existentes
✅ Agregar nuevos campos a las respuestas
✅ Crear validaciones de negocio
✅ Crear endpoints completamente nuevos
✅ Combinar datos de múltiples fuentes
✅ Mejorar manejo de errores
✅ Crear tests automáticos
✅ Mantener estado en los casos de uso
✅ Crear módulos completos desde cero

## 🚀 SIGUIENTES PASOS

Ahora que dominas lo básico, puedes:

1. **Agregar autenticación**: Que solo usuarios registrados puedan usar la API
2. **Agregar paginación**: En los endpoints que devuelven listas
3. **Agregar filtros**: Por ejemplo, `/products?category=electronics`
4. **Conectar una base de datos**: En lugar de APIs externas
5. **Agregar caché**: Para responder más rápido
6. **Agregar documentación Swagger**: Para que otros sepan cómo usar tu API

¡Sigue practicando! 🎉

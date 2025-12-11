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

**PARTE A: Actualizar el contrato del repositorio**

1. Ve a: `internal/domain/user/`
2. Abre el archivo `repository.go`
3. Agrega un nuevo método: `FindAll() ([]*User, error)`

**PARTE B: Implementar la consulta**

4. Ve a: `internal/adapter/repository/`
5. Abre `user_api_repository.go`
6. Crea el método `FindAll` que:
   - Llame a la URL: `{baseURL}/users` (sin ID)
   - Decodifique la respuesta en una lista de usuarios

**PARTE C: Crear la lógica de negocio**

7. Ve a: `internal/usecase/user/`
8. Crea archivo nuevo: `list_users.go`
9. Crea un caso de uso `ListUsersUsecase` que llame al método `FindAll` del repositorio

**PARTE D: Exponer el endpoint**

10. Ve a: `internal/adapter/http/handler/`
11. En `user_handler.go`, agrega el nuevo caso de uso a la estructura
12. Crea un método `List` que llame al caso de uso y devuelva JSON

**PARTE E: Registrar la ruta**

13. Ve a: `internal/infrastructure/http/router.go`
14. Crea el caso de uso de listar usuarios
15. Pásalo al handler
16. Registra la ruta: `r.Get("/users", userHandler.List)`

17. **Prueba**
    - Reinicia el servidor
    - Ejecuta: `curl http://localhost:8080/users`
    - Debes ver una lista de 10 usuarios

### ✅ RESULTADO ESPERADO

- `/users` → Lista completa (array de usuarios)
- `/users/1` → Sigue funcionando (un solo usuario)

### 💡 LO QUE HICISTE

Creaste un endpoint completo nuevo siguiendo todos los pasos de la arquitectura.

---

## EJERCICIO 10 - Endpoint de bienvenida en la raíz

### 📋 LO QUE NECESITAMOS

Cuando alguien entra a `http://localhost:8080/` queremos mostrar un mensaje de bienvenida con información básica:
- Un mensaje amigable
- La versión de la API
- Lista de endpoints disponibles

### 🎯 INSTRUCCIONES

1. **Crea la estructura de datos**
   - Crea carpeta: `internal/domain/welcome/`
   - Crea archivo: `welcome.go`
   - Define estructura con: `Message`, `Version`, `Endpoints` (lista)

2. **Crea la lógica**
   - Crea carpeta: `internal/usecase/welcome/`
   - Crea archivo: `get_welcome.go`
   - Devuelve:
     - Message: "Bienvenido a la API de Ejercicio"
     - Version: "1.1.0"
     - Endpoints: ["/status", "/ping", "/users", "/users/{id}"]

3. **Crea el punto de entrada**
   - En `internal/adapter/http/handler/`
   - Crea: `welcome_handler.go`
   - Método que devuelva la información en JSON

4. **Registra la ruta**
   - En el router, registra: `r.Get("/", welcomeHandler.Get)`

5. **Prueba**
   - `curl http://localhost:8080/` → Debe mostrar la bienvenida

### ✅ RESULTADO ESPERADO

```json
{
  "message": "Bienvenido a la API de Ejercicio",
  "version": "1.1.0",
  "endpoints": ["/status", "/ping", "/users", "/users/{id}"]
}
```

### 💡 LO QUE HICISTE

Creaste una página de inicio para la API que ayuda a los usuarios a descubrir los endpoints.

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

El equipo frontend hace dos llamadas separadas: una a `/users/1` y otra a `/status`. Para mejorar el rendimiento, necesitamos un nuevo endpoint `/user-info/1` que devuelva ambos datos en una sola respuesta.

### 🎯 INSTRUCCIONES

1. **Crea la estructura combinada**
   - Carpeta: `internal/domain/userinfo/`
   - Archivo: `user_info.go`
   - Campos: `User` y `ServerStatus`

2. **Crea la lógica que combina**
   - Carpeta: `internal/usecase/userinfo/`
   - Archivo: `get_user_info.go`
   - Este caso de uso necesita:
     - El caso de uso de GetUser
     - El caso de uso de GetStatus
   - Llama a ambos y combina los resultados

3. **Crea el handler**
   - En `internal/adapter/http/handler/`
   - Archivo: `user_info_handler.go`
   - Extrae el ID, llama al caso de uso, devuelve JSON

4. **Registra la ruta**
   - En router: `r.Get("/user-info/{id}", userInfoHandler.GetByID)`

5. **Prueba**
   - `curl http://localhost:8080/user-info/1`
   - Debe mostrar usuario + status en una respuesta

### ✅ RESULTADO ESPERADO

```json
{
  "user": { "id": 1, "name": "...", ... },
  "server_status": { "message": "...", "version": "1.1.0", ... }
}
```

### 💡 LO QUE HICISTE

Creaste un endpoint optimizado que reduce el número de peticiones del cliente.

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

- Endpoint para obtener un producto: `/products/1`
- Endpoint para listar productos: `/products`

Un producto tiene: ID, Título, Precio, Descripción, Categoría.

### 🎯 INSTRUCCIONES

**PARTE 1: Estructura de datos**
1. Crea: `internal/domain/product/product.go` con los campos necesarios
2. Crea: `internal/domain/product/repository.go` con los métodos FindByID y FindAll

**PARTE 2: Conexión con API externa**
3. Crea: `internal/adapter/repository/product_api_repository.go`
4. Implementa los métodos usando URL: `https://fakestoreapi.com`

**PARTE 3: Lógica de negocio**
5. Crea: `internal/usecase/product/get_product.go` (valida ID entre 1 y 20)
6. Crea: `internal/usecase/product/list_products.go`

**PARTE 4: Endpoints HTTP**
7. Crea: `internal/adapter/http/handler/product_handler.go`
8. Métodos: GetByID y List

**PARTE 5: Registro**
9. En router, registra:
   - `GET /products`
   - `GET /products/{id}`

**PARTE 6: Tests**
10. Crea tests en: `test/usecase/product/`

### ✅ RESULTADO ESPERADO

- `/products` → Lista de productos
- `/products/1` → Un producto específico
- `/products/999` → Error de validación

### 💡 LO QUE HICISTE

Creaste un módulo completo nuevo desde cero, replicando la estructura existente.

---

## 🎓 FELICIDADES

Has completado el taller completo. Ahora puedes:

✅ Cambiar textos y valores en los endpoints
✅ Agregar validaciones de datos
✅ Agregar nuevos campos a las respuestas
✅ Mejorar mensajes de error
✅ Crear endpoints completamente nuevos
✅ Combinar información de múltiples fuentes
✅ Mantener contadores y métricas
✅ Crear módulos completos siguiendo la arquitectura

## 🚀 SIGUIENTES DESAFÍOS

Ahora que dominas lo básico:

1. **Agrega paginación** a los endpoints de listas
2. **Agrega filtros** (ejemplo: `/products?category=electronics`)
3. **Agrega autenticación** para proteger los endpoints
4. **Conecta una base de datos** real en lugar de APIs externas
5. **Agrega documentación automática** con Swagger

¡Sigue practicando! 🎉

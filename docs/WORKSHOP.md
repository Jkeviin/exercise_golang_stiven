# 🎓 WORKSHOP - Clean Architecture en Golang

Este taller te guiará paso a paso para extender la aplicación con nuevas funcionalidades.

**Importante**: Este taller NO incluye código. Solo instrucciones funcionales de qué agregar.

---

## EJERCICIO 1 - Agregar campo "timestamp" al status

### 📋 Objetivo
El endpoint `/status` actualmente devuelve el uptime. Queremos que también devuelve la fecha/hora actual del servidor.

### 🎯 Lo que debes hacer

1. **Agrega un nuevo campo a la entidad Status**
   - Abre el archivo de la entidad Status
   - Agrega un campo llamado `Timestamp` de tipo `string`
   - El tag JSON debe ser `"timestamp"`

2. **Modifica el caso de uso para incluir el timestamp**
   - Abre el archivo del caso de uso de Status
   - En la función que genera el Status, agrega el timestamp actual
   - Usa `time.Now().Format(time.RFC3339)` para el formato

3. **Prueba el cambio**
   - Ejecuta el servidor
   - Llama a `/status`
   - Verifica que ahora aparece el campo `timestamp`

### ✅ Resultado esperado

Cuando llames a `/status`, deberás ver:
```json
{
  "message": "...",
  "version": "1.0.0",
  "uptime": 42,
  "timestamp": "2025-12-11T09:30:45Z"
}
```

### 💡 Errores comunes
- Si el campo no aparece: Verifica que el tag JSON esté correcto
- Si el formato es extraño: Asegúrate de usar `time.RFC3339`

---

## EJERCICIO 2 - Crear endpoint que liste todos los usuarios

### 📋 Objetivo
Actualmente solo puedes obtener un usuario por ID. Queremos un nuevo endpoint `/users` que devuelva la lista completa.

### 🎯 Lo que debes hacer

1. **Agrega el método a la interfaz del repositorio**
   - Abre la interfaz `Repository` de User
   - Agrega un método `FindAll()` que devuelva `([]*User, error)`

2. **Implementa el método en el repositorio API**
   - Abre la implementación del repositorio (UserAPIRepository)
   - Crea la función `FindAll` que:
     - Llame a `{baseURL}/users`
     - Decodifique el JSON en un slice de usuarios
     - Devuelva la lista completa

3. **Crea el nuevo caso de uso**
   - En la carpeta de usecases de user, crea un nuevo archivo
   - Crea el usecase `ListUsersUsecase`
   - El método debe ser `Execute()` y devolver la lista de usuarios

4. **Crea el handler**
   - En la carpeta de handlers, agrega un nuevo método al `UserHandler`
   - Llámalo `List`
   - Debe llamar al nuevo usecase y devolver JSON

5. **Registra la ruta**
   - En el router, agrega la ruta `GET /users`
   - Debe apuntar al nuevo handler

### ✅ Resultado esperado

Cuando llames a `/users`, deberás ver:
```json
[
  {
    "id": 1,
    "name": "Leanne Graham",
    ...
  },
  {
    "id": 2,
    "name": "Ervin Howell",
    ...
  },
  ...
]
```

### 💡 Pistas
- Usa el mismo patrón que `GetByID` pero sin parámetro
- La URL de la API es `/users` (sin el ID)
- El JSON viene como array, no como objeto individual

---

## EJERCICIO 3 - Agregar validación de rango al ID de usuario

### 📋 Objetivo
Actualmente solo validamos que el ID sea mayor que 0. Queremos limitar a IDs entre 1 y 10.

### 🎯 Lo que debes hacer

1. **Modifica la validación en el caso de uso**
   - Abre el usecase `GetUserUsecase`
   - Después de validar que el ID es mayor que 0
   - Agrega una nueva validación: si el ID es mayor que 10, devuelve error
   - El mensaje de error debe ser claro: "el ID debe estar entre 1 y 10"

2. **Prueba el cambio**
   - Llama a `/users/5` → Debe funcionar
   - Llama a `/users/99` → Debe dar error
   - Llama a `/users/0` → Debe dar error

### ✅ Resultado esperado

- `/users/1` a `/users/10` → ✅ Funcionan
- `/users/11` en adelante → ❌ Error "el ID debe estar entre 1 y 10"
- `/users/0` o negativo → ❌ Error "el ID debe ser mayor que 0"

### 💡 Errores comunes
- No olvides que las validaciones van en el **caso de uso**, no en el handler

---

## EJERCICIO 4 - Crear endpoint combinado `/user-with-status/{id}`

### 📋 Objetivo
Crear un endpoint que devuelva la información del usuario Y el status del servidor en una sola respuesta.

### 🎯 Lo que debes hacer

1. **Crea una nueva entidad combinada**
   - En domain, crea un nuevo paquete o archivo
   - Define `UserWithStatus` que tenga:
     - Campo `User` de tipo `user.User`
     - Campo `Status` de tipo `status.Status`
   - Los tags JSON deben ser `"user"` y `"status"`

2. **Crea el nuevo caso de uso**
   - En la carpeta de usecases, crea `GetUserWithStatusUsecase`
   - Este usecase necesita DOS dependencias:
     - El usecase de GetUser
     - El usecase de GetStatus
   - El método `Execute(id int)` debe:
     - Llamar a GetUser con el ID
     - Llamar a GetStatus
     - Combinar ambos resultados en UserWithStatus

3. **Crea el handler**
   - Crea un nuevo handler (o agrégalo al existente)
   - Extrae el ID de la URL
   - Llama al nuevo usecase
   - Devuelve el JSON combinado

4. **Registra la ruta**
   - En el router, agrega `GET /user-with-status/{id}`
   - Conecta con el nuevo handler

### ✅ Resultado esperado

Cuando llames a `/user-with-status/1`, deberás ver:
```json
{
  "user": {
    "id": 1,
    "name": "Leanne Graham",
    ...
  },
  "status": {
    "message": "...",
    "version": "1.0.0",
    "uptime": 42,
    "timestamp": "..."
  }
}
```

### 💡 Concepto importante
Este es un ejemplo de **composición de usecases**. Un usecase puede usar otros usecases para crear funcionalidades más complejas.

---

## EJERCICIO 5 - Agregar un test para el nuevo campo timestamp

### 📋 Objetivo
Crear un test que verifique que el campo timestamp se está generando correctamente.

### 🎯 Lo que debes hacer

1. **Abre el archivo de tests del status usecase**
   - Encuentra el test existente de Status

2. **Agrega una nueva verificación**
   - Después de las verificaciones existentes
   - Verifica que `status.Timestamp` no esté vacío
   - Verifica que tenga un formato válido (contiene "T" y "Z")

3. **Ejecuta los tests**
   - Corre `go test ./...`
   - Verifica que todos pasen

### ✅ Resultado esperado

El test debe pasar y verificar que:
- ✅ El timestamp existe
- ✅ Tiene formato ISO 8601

### 💡 Pistas
- Los tests están en la carpeta `test/usecase/`
- Usa `strings.Contains` para verificar el formato

---

## EJERCICIO 6 - Agregar manejo de errores para API externa

### 📋 Objetivo
Actualmente si la API externa falla, el error no es muy claro. Queremos mejorar los mensajes de error.

### 🎯 Lo que debes hacer

1. **Mejora el manejo de errores en el repositorio**
   - Abre `UserAPIRepository`
   - En el método `FindByID`:
     - Si el código HTTP es 404: Devuelve "usuario no encontrado"
     - Si es 500-599: Devuelve "error en el servidor externo"
     - Si es otro: Devuelve el código específico

2. **Prueba los casos de error**
   - Llama a `/users/999` (no existe)
   - Verifica que el mensaje sea claro

### ✅ Resultado esperado

- `/users/999` → Error claro: "usuario no encontrado"
- API caída → Error: "error en el servidor externo"

---

## EJERCICIO 7 - Agregar logging de peticiones

### 📋 Objetivo
Queremos ver en la consola cada vez que se llama a un endpoint.

### 🎯 Lo que debes hacer

1. **Revisa el middleware de logging**
   - En el router, ya hay un middleware de chi que hace logging
   - Verifica que `middleware.Logger` esté activo

2. **Prueba el logging**
   - Ejecuta el servidor
   - Llama a varios endpoints
   - Observa en la consola que aparecen los logs

### ✅ Resultado esperado

En la consola deberías ver cada petición:
```
GET /status 200
GET /users/1 200
GET /users/999 404
```

---

## EJERCICIO FINAL - Crear módulo completo de productos

### 📋 Objetivo
Crear un módulo completo desde cero para manejar productos, consultando `https://fakestoreapi.com`

### 🎯 Lo que debes hacer

Vas a replicar toda la estructura de Users, pero para Products:

1. **Dominio**
   - Crea `domain/product/product.go`
   - Campos: ID, Title, Price, Description
   - Crea `domain/product/repository.go` con la interfaz

2. **Repositorio**
   - Crea `adapter/repository/product_api_repository.go`
   - URL base: `https://fakestoreapi.com`
   - Implementa `FindByID(id int)`

3. **Caso de uso**
   - Crea `usecase/product/get_product.go`
   - Valida que el ID esté entre 1 y 20
   - Usa el repositorio para obtener el producto

4. **Handler**
   - Crea `adapter/http/handler/product_handler.go`
   - Método `GetByID`
   - Extrae el ID de la URL y llama al usecase

5. **Router**
   - Registra `GET /products/{id}`
   - Conecta con el handler

6. **Tests**
   - Crea `test/usecase/product/get_product_test.go`
   - Usa un mock del repositorio
   - Verifica que el usecase funciona

### ✅ Resultado esperado

Cuando llames a `/products/1`, deberás ver:
```json
{
  "id": 1,
  "title": "Fjallraven - Foldsack No. 1 Backpack",
  "price": 109.95,
  "description": "Your perfect pack..."
}
```

### 💡 Pistas
- Copia la estructura exacta de Users
- Cambia solo los nombres y la URL
- Mantén el mismo patrón de validación

---

## 🎓 Conceptos Aprendidos

Al completar este workshop, habrás practicado:

✅ **Clean Architecture**
- Separación de capas
- Inyección de dependencias
- Interfaces para desacoplamiento

✅ **DDD**
- Entidades de dominio
- Repositorios como abstracciones
- Casos de uso como lógica de aplicación

✅ **Buenas Prácticas**
- Validaciones en el lugar correcto
- Manejo de errores descriptivos
- Tests con mocks
- Composición de usecases

✅ **Escalabilidad**
- Estructura modular
- Fácil agregar nuevos módulos
- Código replicable

---

## 🚀 Siguientes Pasos

Después de dominar estos ejercicios:

1. Agrega autenticación (JWT)
2. Conecta una base de datos real
3. Implementa cache con Redis
4. Agrega paginación a las listas
5. Dockeriza la aplicación

¡Sigue practicando! 🎉

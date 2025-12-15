# 🎯 GUÍA SIMPLE - Entendiendo el Proyecto

> **Para personas que nunca han trabajado con DDD o Clean Architecture**

Esta guía te explica **en lenguaje simple** cómo funciona este proyecto y por qué está organizado así.

---

## 🤔 ¿Por qué tantas carpetas?

Imagina que estás construyendo una casa. No mezclas los cables eléctricos con las tuberías de agua, ¿verdad? 

**Lo mismo pasa aquí**: separamos el código en carpetas según su **responsabilidad**.

---

## 📦 Las 4 Capas (de adentro hacia afuera)

### 1️⃣ **DOMAIN** - "¿Qué es cada cosa?"

📁 `internal/domain/`

**Piensa en esto como un diccionario**:
- Define QUÉ es un Usuario
- Define QUÉ es un Producto
- Define QUÉ campos tiene cada uno

**Ejemplo real**:
```go
// domain/user/user.go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

**En español**: "Un usuario tiene un ID, un nombre y un email"

**🚫 Lo que NO hace**:
- No sabe de dónde vienen los datos
- No sabe cómo mostrarlos en pantalla
- No sabe nada de HTTP, JSON, bases de datos

**✅ Por qué está separado**:
Si mañana cambias de base de datos, o de API externa, o de framework web... **esta parte NO cambia**.

---

### 2️⃣ **USECASE** - "¿Qué puedo hacer?"

📁 `internal/usecase/`

**Piensa en esto como las acciones que puede hacer tu app**:
- "Obtener un usuario por ID"
- "Listar todos los usuarios"
- "Buscar usuarios por email"

**Ejemplo real**:
```go
// usecase/user/get_user.go
func (uc *GetUserUsecase) Execute(id int) (*user.User, error) {
    // 1. Validar que el ID sea válido
    if id <= 0 {
        return nil, errors.New("ID inválido")
    }
    
    // 2. Pedir los datos al repositorio
    return uc.userRepo.FindByID(id)
}
```

**En español**: 
1. "Primero reviso que el ID tenga sentido"
2. "Luego pido los datos"

**🚫 Lo que NO hace**:
- No sabe si los datos vienen de HTTP, de una app móvil, o de la terminal
- No sabe si los datos están en MySQL, PostgreSQL, o una API externa

**✅ Por qué está separado**:
Puedes usar el mismo caso de uso desde:
- Una página web
- Una app móvil
- Un comando de terminal
- Un test automático

---

### 3️⃣ **ADAPTER** - "¿Cómo conecto con el mundo exterior?"

📁 `internal/adapter/`

**Tiene 2 partes**:

#### A) **Handlers** - Reciben peticiones HTTP
📁 `adapter/http/handler/`

```go
// handler/user_handler.go
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    // 1. Extraer el ID de la URL
    idParam := chi.URLParam(r, "id")
    
    // 2. Convertir a número
    id, err := strconv.Atoi(idParam)
    
    // 3. Llamar al caso de uso
    user, err := h.getUserUC.Execute(id)
    
    // 4. Devolver JSON
    json.NewEncoder(w).Encode(user)
}
```

**En español**:
1. "Alguien llamó a `/users/5`"
2. "Extraigo el `5`"
3. "Le pregunto al caso de uso: dame el usuario 5"
4. "Devuelvo la respuesta en formato JSON"

#### B) **Repositories** - Obtienen datos de APIs externas
📁 `adapter/repository/`

```go
// repository/user_api_repository.go
func (r *UserAPIRepository) FindByID(id int) (*user.User, error) {
    // 1. Llamar a la API externa
    url := fmt.Sprintf("%s/users/%d", r.baseURL, id)
    resp, err := http.Get(url)
    
    // 2. Parsear el JSON
    var u user.User
    json.NewDecoder(resp.Body).Decode(&u)
    
    // 3. Devolver el usuario
    return &u, nil
}
```

**En español**:
1. "Voy a llamar a `https://jsonplaceholder.typicode.com/users/5`"
2. "Convierto el JSON que me devuelve en un objeto User"
3. "Devuelvo el usuario"

**✅ Por qué está separado**:
- Si mañana cambias de API externa a base de datos MySQL, **solo cambias este archivo**
- Si cambias de HTTP a GraphQL, **solo cambias los handlers**

---

### 4️⃣ **INFRASTRUCTURE** - "¿Cómo arranco todo?"

📁 `internal/infrastructure/`

**Piensa en esto como el "pegamento"** que conecta todo:

```go
// infrastructure/http/router.go
func SetupRouter(userHandler, statusHandler, pingHandler, welcomeHandler) {
    r := chi.NewRouter()
    
    // Conectar rutas con handlers
    r.Get("/", welcomeHandler.Welcome)
    r.Get("/users/{id}", userHandler.GetByID)
    r.Get("/users", userHandler.GetAll)
    
    return r
}
```

**En español**:
- "Cuando alguien llame a `/users/5`, usa este handler"
- "Cuando alguien llame a `/users`, usa este otro handler"

---

## 🔄 ¿Cómo fluye una petición?

Imagina que alguien llama a: `GET /users/5`

```
1. 🌐 LLEGA LA PETICIÓN HTTP
   "Dame el usuario 5"
   
2. 🚪 ROUTER (infrastructure/http/router.go)
   "Esta petición va al UserHandler"
   
3. 📨 HANDLER (adapter/http/handler/user_handler.go)
   "Extraigo el 5 de la URL"
   "Llamo al caso de uso: getUserUsecase.Execute(5)"
   
4. 🧠 CASO DE USO (usecase/user/get_user.go)
   "¿El 5 es válido? Sí"
   "Llamo al repositorio: userRepo.FindByID(5)"
   
5. 🔌 REPOSITORIO (adapter/repository/user_api_repository.go)
   "Llamo a la API externa"
   "GET https://jsonplaceholder.typicode.com/users/5"
   "Convierto el JSON a un objeto User"
   "Devuelvo el User"
   
6. 🧠 CASO DE USO
   "Recibo el User del repositorio"
   "Lo devuelvo al handler"
   
7. 📨 HANDLER
   "Recibo el User del caso de uso"
   "Lo convierto a JSON"
   "Lo devuelvo como respuesta HTTP"
   
8. 🌐 RESPUESTA SALE
   {"id": 5, "name": "...", "email": "..."}
```

---

## 🎯 ¿Por qué es útil esta separación?

### Ejemplo 1: Cambiar de API a Base de Datos

**Sin separación** (todo mezclado):
- Tienes que cambiar 50 archivos ❌
- Alto riesgo de romper algo ❌

**Con separación** (este proyecto):
- Solo cambias 1 archivo: `adapter/repository/user_api_repository.go` ✅
- El resto sigue funcionando igual ✅

### Ejemplo 2: Agregar una app móvil

**Sin separación**:
- Tienes que reescribir toda la lógica ❌

**Con separación**:
- Reutilizas los casos de uso ✅
- Solo creas nuevos handlers para móvil ✅

### Ejemplo 3: Hacer tests

**Sin separación**:
- Tienes que levantar un servidor HTTP para testear ❌

**Con separación**:
- Testeas el caso de uso directamente ✅
- No necesitas servidor, base de datos, nada ✅

---

## 📚 Glosario de términos

| Término | ¿Qué significa? | Ejemplo en este proyecto |
|---------|----------------|--------------------------|
| **Entidad** | Un "objeto" del negocio | User, Product, Status |
| **Caso de Uso** | Una acción que puede hacer la app | "Obtener usuario", "Listar productos" |
| **Repositorio** | Donde se obtienen los datos | Llamada a API externa |
| **Handler** | Recibe peticiones HTTP | `UserHandler.GetByID()` |
| **Router** | Conecta URLs con handlers | `/users/{id}` → UserHandler |
| **Dominio** | El "diccionario" de la app | Define qué es un User |

---

## 🎓 Reglas de Oro

### ✅ DO (Hacer)
1. **Empieza siempre por la ruta** cuando crees algo nuevo
2. **Valida en 2 lugares**: formato en handler, lógica en caso de uso
3. **Reutiliza casos de uso** cuando puedas
4. **Mantén el dominio limpio**: sin imports de HTTP, JSON, etc.

### ❌ DON'T (No hacer)
1. **No mezcles responsabilidades**: no pongas lógica de negocio en handlers
2. **No hagas que el dominio dependa de nada**: debe ser independiente
3. **No repitas código**: si ves código duplicado, extráelo
4. **No ignores errores**: siempre maneja los errores apropiadamente

---

## 🚀 Para empezar a programar

1. **Lee el ejercicio completo** en `docs/WORKSHOP.md`
2. **Sigue el orden**: Ruta → Handler → Caso de Uso → Repositorio
3. **Mira ejemplos existentes**: compara con código que ya funciona
4. **Prueba frecuentemente**: reinicia el servidor y prueba con `curl`

---

## 💡 Analogía Final

Piensa en este proyecto como un **restaurante**:

- **DOMAIN** = El menú (define qué platos existen)
- **USECASE** = El chef (sabe cómo preparar cada plato)
- **ADAPTER/Handler** = El mesero (toma pedidos de los clientes)
- **ADAPTER/Repository** = El almacén (trae ingredientes)
- **INFRASTRUCTURE** = La puerta y las mesas (organiza todo)

**El cliente** (navegador web) hace un pedido → **el mesero** lo recibe → **el chef** lo prepara usando **ingredientes del almacén** → **el mesero** sirve el plato.

Si mañana cambias de proveedor de ingredientes (API externa → base de datos), **solo cambias el almacén**. El chef, el mesero y el menú siguen igual.

---

## 🎯 Siguiente paso

Ve a `docs/WORKSHOP.md` y empieza con el **Ejercicio 1**. Es súper simple y te ayudará a familiarizarte con la estructura.

¡Éxito! 🚀


# 🏗️ ARQUITECTURA DEL PROYECTO

Este documento explica cómo está construido el proyecto, la responsabilidad de cada carpeta y los principios arquitectónicos aplicados.

---

## 📐 PRINCIPIOS ARQUITECTÓNICOS

Este proyecto sigue dos paradigmas de diseño fundamentales:

### 1. **Clean Architecture (Arquitectura Limpia)**

Organiza el código en capas concéntricas donde:
- **Las capas internas NO conocen a las externas**
- Las dependencias siempre apuntan hacia adentro (hacia el dominio)
- El dominio es el centro y no depende de nada externo

### 2. **Domain-Driven Design (DDD)**

Organiza el código alrededor de los conceptos del negocio:
- **El dominio es el rey**: Todo gira en torno a las entidades y reglas de negocio
- Cada módulo representa un contexto del negocio (User, Ping, Status)
- Las interfaces del dominio definen los contratos, las implementaciones están fuera

---

## 🗂️ ESTRUCTURA DE CARPETAS

```
ejercicio-api/
├── cmd/                    # Punto de entrada de la aplicación
├── internal/               # Código privado de la aplicación
│   ├── domain/            # CAPA 1: Dominio (núcleo del negocio)
│   ├── usecase/           # CAPA 2: Casos de uso (lógica de aplicación)
│   ├── adapter/           # CAPA 3: Adaptadores (implementaciones)
│   ├── infrastructure/    # CAPA 4: Infraestructura (frameworks)
│   └── config/            # Configuración transversal
├── test/                  # Tests organizados por capa
└── docs/                  # Documentación
```

---

## 🎯 RESPONSABILIDAD DE CADA CAPA

### **CAPA 1: Domain (`internal/domain/`)**

> **"El corazón del negocio"**

**Responsabilidad**: Define QUÉ ES cada concepto del negocio, sin saber CÓMO se implementa.

**Contiene**:
- **Entidades**: Estructuras de datos del negocio (`User`, `Status`, `Ping`)
- **Interfaces de Repositorio**: Contratos que definen CÓMO se debe acceder a los datos (sin implementación)
- **Reglas de negocio**: Validaciones y lógica fundamental del dominio

**Características**:
- ✅ **NO depende de ninguna otra capa**
- ✅ **NO importa librerías externas** (excepto las estándar de Go)
- ✅ **NO sabe nada de HTTP, bases de datos o APIs externas**
- ✅ Es la capa más estable, la que menos cambia

**Ejemplo en el proyecto**:
```
internal/domain/
├── user/
│   ├── user.go          # Entidad User (estructura de datos)
│   └── repository.go    # Interface que define CÓMO obtener usuarios
├── status/
│   └── status.go        # Entidad Status
└── ping/
    └── ping.go          # Entidad Ping
```

**Código real**:
- `domain/user/repository.go` define: "necesito un método GetByID que reciba un ID y devuelva un User"
- NO dice: "voy a llamar a una API REST" o "voy a consultar MySQL"

---

### **CAPA 2: Usecase (`internal/usecase/`)**

> **"La lógica de la aplicación"**

**Responsabilidad**: Define CÓMO se ejecutan las operaciones del negocio, orquestando el dominio.

**Contiene**:
- **Casos de uso**: Cada acción que la aplicación puede realizar
- **Lógica de aplicación**: Coordina entidades y repositorios del dominio
- **Reglas de negocio complejas**: Validaciones que involucran múltiples entidades

**Características**:
- ✅ **Depende SOLO del dominio** (`internal/domain`)
- ✅ **NO sabe de HTTP, JSON, o bases de datos**
- ✅ Recibe dependencias por constructor (inyección de dependencias)
- ✅ Fácil de testear (se le inyectan mocks)

**Ejemplo en el proyecto**:
```
internal/usecase/
├── user/
│   └── get_user.go      # Caso de uso: obtener un usuario por ID
├── status/
│   └── get_status.go    # Caso de uso: obtener estado del servidor
└── ping/
    └── ping.go          # Caso de uso: verificar que la API responde
```

**Código real**:
- `usecase/user/get_user.go` dice: "recibo un ID, llamo al repository del dominio, y devuelvo el User"
- NO dice: "parseo el JSON" o "llamo a http.Get"

---

### **CAPA 3: Adapter (`internal/adapter/`)**

> **"Los traductores"**

**Responsabilidad**: Implementa las interfaces del dominio adaptando tecnologías externas.

**Contiene**:
- **Implementaciones de Repositorios**: Código real que obtiene datos (de APIs, DB, archivos)
- **Handlers HTTP**: Reciben requests HTTP y llaman a los casos de uso
- **Adaptadores de datos**: Convierten formatos externos al formato del dominio

**Características**:
- ✅ **Implementa interfaces definidas en el dominio**
- ✅ **Conoce tecnologías externas** (HTTP clients, ORMs, etc.)
- ✅ **Traduce** entre el mundo externo y el dominio
- ✅ Es la capa más cambiante (si cambias de API o DB, solo tocas esto)

**Ejemplo en el proyecto**:
```
internal/adapter/
├── repository/
│   └── user_api_repository.go    # Implementa user.Repository llamando a JSONPlaceholder
└── http/
    └── handler/
        ├── user_handler.go        # Recibe HTTP request, llama usecase, devuelve HTTP response
        ├── status_handler.go
        └── ping_handler.go
```

**Código real**:
- `adapter/repository/user_api_repository.go` implementa la interface `domain/user/repository.go`
- Aquí SÍ se usa `http.Get()`, se parsea JSON, se manejan errores HTTP
- `adapter/http/handler/user_handler.go` recibe el `http.ResponseWriter`, extrae parámetros de la URL, llama al usecase, y escribe JSON

---

### **CAPA 4: Infrastructure (`internal/infrastructure/`)**

> **"La fontanería"**

**Responsabilidad**: Configuración de frameworks y herramientas externas.

**Contiene**:
- **Router HTTP**: Configuración del servidor web y rutas
- **Middleware**: Logging, CORS, autenticación
- **Configuración de librerías**: Setup de ORMs, clientes HTTP, etc.

**Características**:
- ✅ **Ensambla todas las capas**
- ✅ **Configura frameworks** (chi router, middleware, etc.)
- ✅ **NO contiene lógica de negocio**
- ✅ Fácil de reemplazar (puedes cambiar chi por gin sin tocar el dominio)

**Ejemplo en el proyecto**:
```
internal/infrastructure/
└── http/
    └── router.go    # Configura chi router, registra rutas, aplica middleware
```

**Código real**:
- `infrastructure/http/router.go` crea el router de chi, define las rutas (`/ping`, `/users/{id}`), y conecta cada ruta con su handler

---

### **Config (`internal/config/`)**

> **"Configuración centralizada"**

**Responsabilidad**: Maneja configuración de la aplicación (puertos, URLs, timeouts).

**Características**:
- ✅ **Transversal**: Todas las capas pueden usarlo
- ✅ Lee variables de entorno
- ✅ Define valores por defecto

---

### **CMD (`cmd/`)**

> **"El punto de entrada"**

**Responsabilidad**: Inicializa y arranca la aplicación.

**Contiene**:
- `main.go`: Crea todas las dependencias y arranca el servidor

**Código real**:
```
1. Carga configuración
2. Crea repositorios (adapters)
3. Crea casos de uso (inyectando repositorios)
4. Crea handlers (inyectando casos de uso)
5. Configura router (infrastructure)
6. Arranca servidor HTTP
```

---

### **Test (`test/`)**

> **"Pruebas organizadas"**

**Responsabilidad**: Tests unitarios e integración.

**Estructura**:
```
test/
└── usecase/
    ├── user/
    ├── status/
    └── ping/
```

**Características**:
- ✅ Organizados por capa y dominio
- ✅ Usan mocks para aislar dependencias
- ✅ Siguen el mismo patrón de estructura que el código

---

## 🔄 FLUJO DE DEPENDENCIAS

```
┌─────────────────────────────────────────────┐
│          cmd/app/main.go                    │  Punto de entrada
│  (Inicializa todo y conecta las capas)     │
└─────────────────┬───────────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────────┐
│    internal/infrastructure/                 │  CAPA 4: Infraestructura
│  - Configura frameworks (chi router)        │  (Frameworks y herramientas)
│  - Registra rutas                           │
└─────────────────┬───────────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────────┐
│    internal/adapter/                        │  CAPA 3: Adaptadores
│  - HTTP Handlers (reciben requests)         │  (Implementaciones concretas)
│  - Repositories (llaman a APIs externas)    │
└─────────────────┬───────────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────────┐
│    internal/usecase/                        │  CAPA 2: Casos de Uso
│  - Orquesta la lógica de negocio            │  (Lógica de aplicación)
│  - Coordina repositorios y entidades        │
└─────────────────┬───────────────────────────┘
                  │
                  ↓
┌─────────────────────────────────────────────┐
│    internal/domain/                         │  CAPA 1: Dominio
│  - Entidades de negocio                     │  (Núcleo del negocio)
│  - Interfaces de repositorio                │  NO DEPENDE DE NADA
│  - Reglas de negocio fundamentales          │
└─────────────────────────────────────────────┘
```

**Regla de Oro**: **Las flechas siempre apuntan hacia abajo (hacia el dominio)**.

---

## 🎬 FLUJO DE UNA REQUEST HTTP

Ejemplo: `GET /users/1`

```
1. REQUEST LLEGA
   ↓
2. infrastructure/http/router.go
   - Chi router recibe el request
   - Extrae el parámetro {id}
   - Llama al handler correspondiente
   ↓
3. adapter/http/handler/user_handler.go
   - Valida que el ID sea válido
   - Llama al caso de uso GetUser
   ↓
4. usecase/user/get_user.go
   - Recibe el ID
   - Llama al repository (interface del dominio)
   ↓
5. adapter/repository/user_api_repository.go
   - Implementa la interface
   - Hace HTTP GET a JSONPlaceholder API
   - Parsea el JSON
   - Devuelve una entidad User (del dominio)
   ↓
6. usecase/user/get_user.go
   - Recibe el User
   - Lo devuelve al handler
   ↓
7. adapter/http/handler/user_handler.go
   - Convierte User a JSON
   - Escribe la respuesta HTTP
   ↓
8. RESPONSE SALE
```

---

## ✅ VENTAJAS DE ESTA ARQUITECTURA

### **1. Separación de Responsabilidades**
- Cada capa tiene un propósito claro
- Es fácil encontrar dónde hacer cambios

### **2. Testeable**
- Los casos de uso no dependen de HTTP o bases de datos
- Se pueden testear con mocks fácilmente

### **3. Independiente de Frameworks**
- Puedes cambiar chi por gin, echo o net/http sin tocar el dominio
- Puedes cambiar la API externa sin tocar los casos de uso

### **4. Escalable**
- Agregar nuevos dominios es copiar la estructura
- Cada dominio está aislado en su carpeta

### **5. Mantenible**
- La lógica de negocio está en el dominio, no mezclada con HTTP o SQL
- Los cambios en una capa no afectan a las demás

---

## 📚 EJEMPLO PRÁCTICO: Dominio USER

### **domain/user/user.go** (Entidad)
```go
type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Username string `json:"username"`
}
```
→ Define QUÉ es un User

### **domain/user/repository.go** (Interface)
```go
type Repository interface {
    GetByID(id int) (*User, error)
}
```
→ Define CÓMO se debe obtener un User (contrato)

### **usecase/user/get_user.go** (Caso de Uso)
```go
type GetUserUsecase struct {
    userRepo user.Repository  // Depende de la interface del dominio
}

func (uc *GetUserUsecase) Execute(id int) (*user.User, error) {
    return uc.userRepo.GetByID(id)  // Llama a la interface
}
```
→ Orquesta: recibe ID, llama al repositorio, devuelve User

### **adapter/repository/user_api_repository.go** (Implementación)
```go
type UserAPIRepository struct {
    baseURL string
}

func (r *UserAPIRepository) GetByID(id int) (*user.User, error) {
    resp, err := http.Get(fmt.Sprintf("%s/users/%d", r.baseURL, id))
    // ... parsea JSON, maneja errores ...
    return &user.User{...}, nil
}
```
→ Implementa la interface: hace HTTP GET, parsea JSON, devuelve User

### **adapter/http/handler/user_handler.go** (Handler)
```go
type UserHandler struct {
    getUserUsecase *usecase.GetUserUsecase
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")  // Extrae parámetro de URL
    user, err := h.getUserUsecase.Execute(id)  // Llama al caso de uso
    json.NewEncoder(w).Encode(user)  // Devuelve JSON
}
```
→ Traduce HTTP a dominio y viceversa

---

## 🎓 RESUMEN

| Capa | Responsabilidad | Depende de | Tecnologías |
|------|----------------|------------|-------------|
| **Domain** | Definir el negocio | Nada | Solo Go estándar |
| **Usecase** | Lógica de aplicación | Domain | Solo Go estándar |
| **Adapter** | Implementar interfaces | Domain, Usecase | HTTP, JSON, APIs |
| **Infrastructure** | Configurar frameworks | Adapter, Usecase | chi, middleware |
| **Config** | Configuración | - | Variables de entorno |
| **CMD** | Inicializar app | Todas | Todas |

---

## 🔍 CÓMO AGREGAR NUEVA FUNCIONALIDAD

**Ejemplo**: Agregar endpoint `GET /products/{id}`

### 1️⃣ Domain
```
internal/domain/product/
├── product.go       # Entidad Product
└── repository.go    # Interface Repository
```

### 2️⃣ Usecase
```
internal/usecase/product/
└── get_product.go   # Caso de uso GetProduct
```

### 3️⃣ Adapter
```
internal/adapter/
├── repository/
│   └── product_api_repository.go  # Implementa interface
└── http/handler/
    └── product_handler.go          # Handler HTTP
```

### 4️⃣ Infrastructure
Registrar ruta en `internal/infrastructure/http/router.go`

### 5️⃣ CMD
Conectar todo en `cmd/app/main.go`

---

## 📖 REFERENCIAS

- **Clean Architecture**: Robert C. Martin (Uncle Bob)
- **Domain-Driven Design**: Eric Evans
- **Go Project Layout**: https://github.com/golang-standards/project-layout

---

## 💡 CONCLUSIÓN

Esta arquitectura puede parecer "mucha carpeta" para un proyecto simple, pero:

✅ **Es escalable**: Agregar 100 dominios más es fácil
✅ **Es mantenible**: Cada cosa está en su lugar
✅ **Es profesional**: Así se construyen sistemas reales
✅ **Es educativa**: Aprendes patrones que aplican a cualquier lenguaje

**La clave**: Respetar el flujo de dependencias (siempre hacia el dominio) y no mezclar responsabilidades.

# 🔥 Hot Reload con Air

Esta guía explica cómo usar **Air** para desarrollo con recarga automática.

---

## ❓ ¿Qué es Hot Reload?

**Hot Reload** (o "recarga en caliente") significa que el servidor **se reinicia automáticamente** cada vez que guardas cambios en tu código.

**Sin Hot Reload**:
1. Haces cambios en el código
2. Detienes el servidor (Ctrl+C)
3. Vuelves a ejecutar `go run cmd/app/main.go`
4. Pruebas los cambios
5. Repites desde el paso 1 🔄😫

**Con Hot Reload (Air)**:
1. Haces cambios en el código
2. Guardas el archivo (Ctrl+S)
3. ✅ El servidor se reinicia automáticamente
4. Pruebas los cambios
5. Repites desde el paso 1 🚀😎

---

## 📥 Instalación

### Instalar Air (solo una vez)

```bash
go install github.com/air-verse/air@latest
```

Esto instala Air en `~/go/bin/air`

### Verificar instalación

```bash
air -v
```

Si te da error "command not found", agrega Go bin a tu PATH:

```bash
# En ~/.zshrc o ~/.bashrc
export PATH=$PATH:$HOME/go/bin
```

---

## 🚀 Uso

### Opción 1: Comando directo

```bash
air
```

### Opción 2: Con Make (recomendado)

```bash
make dev
```

---

## 🎨 Salida de Air

Cuando ejecutas `air`, verás algo así:

```
  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.63.4, built with Go go1.25.5

watching .
!exclude tmp
!exclude testdata
building...
running...

🚀 Servidor iniciado en http://localhost:8080

```

### Cuando haces un cambio:

```
main.go has changed
building...
running...

🚀 Servidor iniciado en http://localhost:8080
```

✅ **¡Automático!** No tuviste que hacer nada.

---

## ⚙️ Configuración

El proyecto incluye un archivo `.air.toml` con la configuración:

```toml
[build]
  # Comando para compilar
  cmd = "go build -o ./tmp/main ./cmd/app/main.go"
  
  # Binario generado
  bin = "tmp/main"
  
  # Archivos a vigilar
  include_ext = ["go", "tpl", "tmpl", "html"]
  
  # Carpetas a ignorar
  exclude_dir = ["assets", "tmp", "vendor", "testdata", "test"]
  
  # Ignorar archivos de test
  exclude_regex = ["_test.go"]
  
  # Retrasar el reinicio (ms) después de detectar cambios
  delay = 1000
```

### ¿Qué hace cada opción?

- **`cmd`**: Comando para compilar tu app
- **`bin`**: Dónde guarda el ejecutable temporal
- **`include_ext`**: Extensiones de archivos que vigila
- **`exclude_dir`**: Carpetas que ignora
- **`exclude_regex`**: Patrones de archivos a ignorar (como tests)
- **`delay`**: Espera 1 segundo antes de reiniciar (útil si guardas varios archivos seguidos)

---

## 🎯 Casos de Uso

### 1. Desarrollo de nuevos endpoints

```go
// 1. Agregas un nuevo endpoint en router.go
r.Get("/products/{id}", productHandler.GetByID)

// 2. Guardas (Ctrl+S)
// ✅ Air reinicia automáticamente el servidor

// 3. Pruebas inmediatamente
curl http://localhost:8080/products/1
```

### 2. Modificar lógica de negocio

```go
// 1. Cambias un usecase
func (uc *GetUserUsecase) Execute(id int) (*user.User, error) {
    // Agregas validación
    if id <= 0 {
        return nil, errors.New("ID inválido")
    }
    return uc.userRepo.GetByID(id)
}

// 2. Guardas
// ✅ Air reinicia y ya tienes la validación activa
```

### 3. Trabajar con múltiples archivos

```
# Si modificas 3 archivos:
- internal/domain/user/user.go
- internal/usecase/user/get_user.go
- internal/adapter/http/handler/user_handler.go

Air espera 1 segundo (delay) y reinicia UNA VEZ
con todos los cambios aplicados
```

---

## 📂 Archivos y Carpetas

### `.air.toml`
Archivo de configuración de Air (ya incluido en el proyecto)

### `tmp/`
Carpeta donde Air guarda los binarios temporales (ignorada por Git)

### `build-errors.log`
Log de errores de compilación (ignorado por Git)

---

## 🔍 Troubleshooting

### ❌ "command not found: air"

**Solución**: Agrega Go bin a tu PATH

```bash
export PATH=$PATH:$HOME/go/bin
```

### ❌ Air no detecta cambios

**Solución**: Verifica que los archivos estén en las carpetas vigiladas

```bash
# Carpetas vigiladas:
internal/
cmd/

# Carpetas ignoradas:
tmp/
test/
vendor/
```

### ❌ "address already in use"

**Solución**: Otro proceso está usando el puerto 8080

```bash
# Buscar proceso en puerto 8080
lsof -ti:8080

# Matarlo
kill -9 $(lsof -ti:8080)

# O cambiar el puerto en config
export SERVER_PORT=8081
air
```

### ❌ El servidor reinicia muy seguido

**Solución**: Aumenta el `delay` en `.air.toml`

```toml
[build]
  delay = 2000  # Espera 2 segundos en lugar de 1
```

---

## 💡 Tips

### 1. Usar con Git

Air ignora automáticamente:
- Archivos de test (`*_test.go`)
- Carpetas temporales (`tmp/`)
- Archivos no relacionados con código

Así puedes hacer commits sin que Air reinicie constantemente.

### 2. Ver logs limpios

Air limpia la terminal en cada reinicio. Si quieres mantener el historial:

```toml
[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

### 3. Ejecutar tests automáticamente

Si quieres que también ejecute tests:

```toml
[build]
  cmd = "go test ./... && go build -o ./tmp/main ./cmd/app/main.go"
```

---

## 📊 Comparación

| Característica | `go run` | `air` (Hot Reload) |
|----------------|----------|---------------------|
| Reinicio manual | ❌ Sí | ✅ Automático |
| Velocidad | 🐢 Lento (recompila todo) | 🚀 Rápido (solo cambios) |
| Productividad | 😐 Media | 🔥 Alta |
| Setup | ✅ Ninguno | ⚙️ Instalar Air |
| Para producción | ✅ No usar | ❌ Solo desarrollo |

---

## 🎓 Resumen

1. **Instala Air**: `go install github.com/air-verse/air@latest`
2. **Ejecuta**: `air` o `make dev`
3. **Desarrolla**: Haz cambios y guarda
4. **Disfruta**: El servidor se reinicia solo 🎉

**Air es obligatorio para desarrollo profesional en Go**. Te ahorra horas de trabajo.

---

## 🔗 Referencias

- [Air en GitHub](https://github.com/air-verse/air)
- [Documentación oficial](https://github.com/air-verse/air#readme)

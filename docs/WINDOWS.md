# 🪟 Guía para Windows

Esta guía explica cómo usar el proyecto en **Windows**, ya que `make` no está disponible por defecto.

---

## ❓ ¿Por qué no funciona `make` en Windows?

`make` es una herramienta de **Unix/Linux/macOS** que no viene instalada en Windows. Por eso, cuando intentas ejecutar `make dev` o `make run`, obtienes un error:

```
'make' no se reconoce como un comando interno o externo
```

---

## ✅ SOLUCIONES

Hay **3 opciones** para trabajar en Windows:

### **OPCIÓN 1: Scripts .bat (Más Fácil) ⭐**

He creado scripts `.bat` equivalentes a los comandos `make`. Solo ejecuta los `.bat` desde la carpeta `scripts/`:

```cmd
REM Instalar dependencias
scripts\deps.bat

REM Ejecutar con hot reload (recomendado)
scripts\dev.bat

REM Ejecutar sin hot reload
scripts\run.bat

REM Ejecutar tests
scripts\test.bat

REM Compilar ejecutable
scripts\build.bat

REM Ver ayuda
scripts\help.bat
```

**✅ ESTA ES LA FORMA MÁS SIMPLE Y RECOMENDADA**

---

### **OPCIÓN 2: Comandos directos en PowerShell/CMD**

Si prefieres ejecutar los comandos directamente:

```powershell
# Instalar dependencias
go mod tidy
go mod download

# Ejecutar servidor con hot reload
air

# Ejecutar servidor sin hot reload
go run cmd\app\main.go

# Ejecutar tests
go test .\... -v

# Compilar
go build -o ejercicio-api.exe cmd\app\main.go
```

---

### **OPCIÓN 3: Instalar Make en Windows (Avanzado)**

Si quieres usar `make` como en Mac/Linux, tienes estas opciones:

#### A) **Chocolatey** (Package manager para Windows)

```powershell
# 1. Instalar Chocolatey (PowerShell como Administrador)
Set-ExecutionPolicy Bypass -Scope Process -Force
[System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072
iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))

# 2. Instalar Make
choco install make
```

#### B) **Scoop** (Alternativa a Chocolatey)

```powershell
# 1. Instalar Scoop (PowerShell)
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
irm get.scoop.sh | iex

# 2. Instalar Make
scoop install make
```

#### C) **Git Bash** (Viene con Git para Windows)

Si tienes Git instalado, usa **Git Bash** que incluye `make`:

1. Abre **Git Bash** (no CMD ni PowerShell)
2. Ejecuta los comandos `make` normalmente:

```bash
make dev
make run
make test
```

---

## 🔥 Configurar Hot Reload (Air) en Windows

### 1. Instalar Air

```powershell
go install github.com/air-verse/air@latest
```

### 2. Agregar Go bin al PATH

Air se instala en `%USERPROFILE%\go\bin\air.exe`

**Agregar al PATH**:

1. Presiona `Win + R`
2. Escribe `sysdm.cpl` y presiona Enter
3. Ve a la pestaña **Avanzado**
4. Click en **Variables de entorno**
5. En "Variables del sistema", busca **Path** y edítala
6. Agrega: `%USERPROFILE%\go\bin`
7. Click **Aceptar** en todo

### 3. Verificar instalación

Abre una **nueva terminal** (PowerShell o CMD):

```cmd
air -v
```

Deberías ver:

```
  __    _   ___  
 / /\  | | | |_) 
/_/--\ |_| |_| \_ v1.63.4
```

### 4. Ejecutar con Air

```cmd
REM Opción 1: Script
scripts\dev.bat

REM Opción 2: Comando directo
air
```

---

## 📁 Scripts Disponibles

| Script | Equivalente Make | Descripción |
|--------|------------------|-------------|
| `scripts\help.bat` | `make help` | Muestra ayuda |
| `scripts\deps.bat` | `make deps` | Instala dependencias |
| `scripts\dev.bat` | `make dev` | 🔥 Hot reload con Air |
| `scripts\run.bat` | `make run` | Ejecutar sin hot reload |
| `scripts\test.bat` | `make test` | Ejecutar tests |
| `scripts\build.bat` | `make build` | Compilar ejecutable |

---

## 🎯 Flujo de Trabajo Típico (Windows)

### Primera vez:

```cmd
REM 1. Instalar dependencias
scripts\deps.bat

REM 2. Instalar Air (hot reload)
go install github.com/air-verse/air@latest

REM 3. Verificar que todo funciona
scripts\test.bat
```

### Desarrollo diario:

```cmd
REM Ejecutar con hot reload
scripts\dev.bat

REM Haces cambios, guardas (Ctrl+S)
REM ✅ El servidor se reinicia automáticamente
```

---

## 🔧 Diferencias entre Windows y Mac/Linux

| Comando | Mac/Linux | Windows |
|---------|-----------|---------|
| Separador de rutas | `/` | `\` |
| Ejecutar script | `./script.sh` | `script.bat` |
| Variable de entorno | `$HOME` | `%USERPROFILE%` |
| Make | ✅ Instalado | ❌ No disponible |
| Ejecutable | `app` | `app.exe` |

---

## ⚠️ Problemas Comunes

### ❌ "air no se reconoce como comando"

**Causa**: Air no está en el PATH

**Solución**:

```powershell
# Ejecutar directamente con ruta completa
%USERPROFILE%\go\bin\air.exe

# O agregar al PATH (ver sección anterior)
```

### ❌ "go: no such file or directory"

**Causa**: Usas `/` en lugar de `\` en Windows

**Solución**: Los scripts `.bat` ya usan `\` correctamente

```cmd
REM ❌ Incorrecto en Windows
go run cmd/app/main.go

REM ✅ Correcto en Windows
go run cmd\app\main.go

REM ✅ Mejor: Usa el script
scripts\run.bat
```

### ❌ "The term 'make' is not recognized"

**Causa**: Make no está instalado

**Solución**: Usa los scripts `.bat` o instala Make (ver Opción 3)

---

## 🎨 PowerShell vs CMD

Ambos funcionan, pero **PowerShell** es más moderno:

### CMD (Command Prompt):

```cmd
C:\proyecto> scripts\dev.bat
```

### PowerShell:

```powershell
PS C:\proyecto> .\scripts\dev.bat
```

**Recomendación**: Usa **PowerShell** o **Windows Terminal** (Windows 11).

---

## 🚀 Inicio Rápido (Resumen)

**Para empezar en Windows**:

```cmd
REM 1. Instalar dependencias
go mod tidy

REM 2. Instalar Air (opcional, para hot reload)
go install github.com/air-verse/air@latest

REM 3. Ejecutar servidor
scripts\dev.bat

REM 4. Probar endpoints
curl http://localhost:8080/status
curl http://localhost:8080/ping
```

---

## 📚 Más Información

- **Hot Reload**: Lee `docs/HOT_RELOAD.md`
- **Arquitectura**: Lee `ARQUITECTURA.md`
- **Workshop**: Lee `docs/WORKSHOP.md`

---

## ✅ Resumen

| Necesitas | Solución |
|-----------|----------|
| Ejecutar comandos rápido | Usa `scripts\*.bat` ⭐ |
| Hot reload | Instala Air + `scripts\dev.bat` |
| Compatibilidad con `make` | Instala Make o usa Git Bash |
| PowerShell vs CMD | Cualquiera funciona, PowerShell es mejor |

**La forma más simple**: Usa los scripts `.bat` de la carpeta `scripts/`.

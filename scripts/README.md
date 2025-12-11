# 🪟 Scripts para Windows

Esta carpeta contiene scripts `.bat` para usar en **Windows** (equivalentes a los comandos `make`).

---

## 📋 Scripts Disponibles

| Script | Equivalente Make | Descripción |
|--------|------------------|-------------|
| `help.bat` | `make help` | Muestra todos los comandos disponibles |
| `deps.bat` | `make deps` | Instala dependencias del proyecto |
| `dev.bat` | `make dev` | 🔥 Ejecuta servidor con hot reload (Air) |
| `run.bat` | `make run` | Ejecuta el servidor sin hot reload |
| `test.bat` | `make test` | Ejecuta todos los tests |
| `build.bat` | `make build` | Compila el ejecutable |

---

## 🚀 Uso

### Desde CMD o PowerShell:

```cmd
REM Ver ayuda
scripts\help.bat

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
```

---

## 💡 Tip

Para desarrollo, **siempre usa `dev.bat`** (con hot reload). El servidor se reiniciará automáticamente cada vez que guardes cambios.

---

## 📚 Más Información

Lee la **[Guía Completa de Windows](../docs/WINDOWS.md)** para:
- Cómo instalar Air en Windows
- Alternativas a los scripts
- Troubleshooting de problemas comunes
- Cómo instalar `make` en Windows (opcional)

@echo off
REM Script para Windows: Instalar dependencias

echo.
echo 📦 Instalando dependencias...
echo.

go mod tidy
go mod download

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ✅ Dependencias instaladas correctamente
) else (
    echo.
    echo ❌ Error al instalar dependencias
    exit /b 1
)

@echo off
REM Script para Windows: Compilar ejecutable

echo.
echo 🔨 Compilando...
echo.

go build -o ejercicio-api.exe cmd\app\main.go

if %ERRORLEVEL% EQU 0 (
    echo.
    echo ✅ Ejecutable creado: ejercicio-api.exe
) else (
    echo.
    echo ❌ Error en la compilacion
    exit /b 1
)

@echo off
REM Script para Windows: Ejecutar servidor con hot reload (Air)

echo.
echo 🔥 Iniciando servidor con hot reload (Air)...
echo 💡 Los cambios se recargaran automaticamente
echo.

where air >nul 2>nul
if %ERRORLEVEL% NEQ 0 (
    echo ❌ Air no esta instalado.
    echo.
    echo Para instalar Air, ejecuta:
    echo go install github.com/air-verse/air@latest
    echo.
    echo Luego agrega %%USERPROFILE%%\go\bin a tu PATH
    exit /b 1
)

air

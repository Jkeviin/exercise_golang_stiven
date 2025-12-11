# 📮 Guía de Importación en Postman

## 🚀 Cómo importar la colección

### Opción 1: Importar el archivo directamente

1. Abre Postman
2. Click en **"Import"** (esquina superior izquierda)
3. Arrastra el archivo `postman_collection.json` o click en "Upload Files"
4. Click en **"Import"**
5. ¡Listo! Verás la colección "Ejercicio API - Clean Architecture"

### Opción 2: Importar desde código

1. Abre Postman
2. Click en **"Import"**
3. Pega el contenido del archivo `postman_collection.json`
4. Click en **"Import"**

---

## 📋 Contenido de la Colección

### 1. **Status** (Estado del servidor)
- `GET /status` - Información del servidor (versión, uptime, etc.)

### 2. **Health Check** (Verificación)
- `GET /ping` - Respuesta rápida para verificar que el servidor está activo

### 3. **Users** (Usuarios)
- `GET /users/1` - Obtener usuario con ID 1
- `GET /users/2` - Obtener usuario con ID 2
- `GET /users/5` - Obtener usuario con ID 5

### 4. **Error Cases** (Casos de error)
- `GET /users/0` - Validación: ID = 0
- `GET /users/-1` - Validación: ID negativo
- `GET /users/999` - Validación: ID muy grande
- `GET /users/abc` - Validación: ID no numérico

---

## ⚙️ Configuración

### Variable de entorno

La colección usa una variable `{{base_url}}` configurada por defecto en:
```
http://localhost:8080
```

Si tu servidor corre en otro puerto, puedes cambiarla:

1. En Postman, ve a la colección "Ejercicio API"
2. Click en **"Variables"**
3. Cambia el valor de `base_url`
4. Guarda los cambios

**Ejemplos de valores**:
- `http://localhost:8080` (puerto por defecto)
- `http://localhost:9000` (puerto personalizado)
- `https://mi-api-desplegada.com` (producción)

---

## 🧪 Cómo usar la colección

### Paso 1: Inicia tu servidor
```bash
go run cmd/app/main.go
```

### Paso 2: Prueba los endpoints

1. **Status del servidor**:
   - Selecciona "Status" → "Get Server Status"
   - Click en **"Send"**
   - Deberías ver: `{"message":"...","version":"1.0.0","uptime":...}`

2. **Ping (health check)**:
   - Selecciona "Health Check" → "Ping"
   - Click en **"Send"**
   - Deberías ver: `{"message":"pong"}`

3. **Obtener usuario**:
   - Selecciona "Users" → "Get User by ID"
   - Click en **"Send"**
   - Deberías ver información del usuario

4. **Probar validaciones**:
   - Ve a "Error Cases"
   - Prueba cada request
   - Observa los diferentes mensajes de error

---

## 💡 Tips

### Cambiar el ID del usuario dinámicamente

Para el request "Get User by ID":

1. Click en el request
2. Ve a la pestaña **"Params"**
3. Cambia el valor de `id`
4. Click en **"Send"**

### Ver respuestas bonitas

En Postman, después de enviar un request:
- Pestaña **"Body"** → Formato **"Pretty"** → Selecciona **"JSON"**
- Verás el JSON con colores y indentado

### Guardar respuestas de ejemplo

1. Envía un request
2. Click en **"Save Response"** → **"Save as example"**
3. Ponle un nombre descriptivo
4. La próxima vez verás el ejemplo antes de enviar

---

## 📊 Respuestas Esperadas

### GET /status
```json
{
  "message": "La aplicación está funcionando correctamente",
  "version": "1.0.0",
  "uptime": 42
}
```

### GET /ping
```json
{
  "message": "pong"
}
```

### GET /users/1
```json
{
  "id": 1,
  "name": "Leanne Graham",
  "email": "Sincere@april.biz",
  "username": "Bret"
}
```

### GET /users/0 (error)
```
el ID debe ser mayor que 0
```

---

## 🔧 Troubleshooting

### Error: "Could not get response"
- ✅ Verifica que el servidor esté corriendo
- ✅ Verifica que la URL sea `http://localhost:8080`
- ✅ Revisa que el puerto no esté ocupado

### Error: "404 Not Found"
- ✅ Verifica que la ruta sea correcta
- ✅ Asegúrate de incluir la `/` inicial: `/users/1`

### Error: "Connection refused"
- ✅ Inicia el servidor: `go run cmd/app/main.go`
- ✅ Verifica que no haya errores en la consola

---

## 🎯 Próximos Pasos

Después de probar estos endpoints:

1. **Modifica el código** siguiendo el `WORKSHOP.md`
2. **Prueba tus cambios** con Postman
3. **Agrega nuevos endpoints** a la colección
4. **Exporta la colección actualizada** (Collection → Export)

---

## 📝 Notas

- Esta colección se actualizará conforme agregues más endpoints
- Puedes duplicar requests para crear variaciones
- Usa Environments para manejar múltiples servidores (dev, staging, prod)

¡Disfruta probando tu API! 🚀


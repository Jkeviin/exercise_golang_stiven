package main

import (
	"ejercicio-api/internal/adapter/http/handler"
	"ejercicio-api/internal/adapter/repository"
	"ejercicio-api/internal/config"
	httpInfra "ejercicio-api/internal/infrastructure/http"
	pingUsecase "ejercicio-api/internal/usecase/ping"
	statusUsecase "ejercicio-api/internal/usecase/status"
	userUsecase "ejercicio-api/internal/usecase/user"
	"log"
)

func main() {
	// 1️⃣ CARGAR CONFIGURACIÓN
	cfg := config.Load()

	// 2️⃣ CREAR REPOSITORIOS (Capa de Adaptadores)
	// Los repositorios implementan las interfaces definidas en el dominio
	userRepo := repository.NewUserAPIRepository(cfg.ExternalAPIURL)

	// 3️⃣ CREAR CASOS DE USO (Capa de Aplicación)
	// Inyectamos los repositorios en los casos de uso
	getUserUsecase := userUsecase.NewGetUserUsecase(userRepo)
	getStatusUsecase := statusUsecase.NewGetStatusUsecase()
	pingUsecase := pingUsecase.NewPingUsecase()

	// 4️⃣ CREAR HANDLERS HTTP (Capa de Adaptadores)
	// Inyectamos los casos de uso en los handlers
	userHandler := handler.NewUserHandler(getUserUsecase)
	statusHandler := handler.NewStatusHandler(getStatusUsecase)
	pingHandler := handler.NewPingHandler(pingUsecase)

	// 5️⃣ CONFIGURAR ROUTER (Capa de Infraestructura)
	// Conectamos las rutas con los handlers
	router := httpInfra.SetupRouter(userHandler, statusHandler, pingHandler)

	// 6️⃣ INICIAR SERVIDOR HTTP
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", cfg.ServerPort)
	if err := httpInfra.Start(cfg, router); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}

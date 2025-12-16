package main

import (
	"ejercicio-api/internal/config"
	"ejercicio-api/internal/handler"
	"ejercicio-api/internal/repository"
	"ejercicio-api/internal/router"
	pingUC "ejercicio-api/internal/usecase/ping"
	statusUC "ejercicio-api/internal/usecase/status"
	userUC "ejercicio-api/internal/usecase/user"
	welcomeUC "ejercicio-api/internal/usecase/welcome"
	"log"
)

func main() {
	// 1️⃣ CARGAR CONFIGURACIÓN
	cfg := config.Load()

	// 2️⃣ CREAR REPOSITORIOS
	userRepo := repository.NewUserRepository(cfg.ExternalAPIURL)

	// 3️⃣ CREAR CASOS DE USO
	getUserUsecase := userUC.NewGetUserUsecase(userRepo)
	getUserAllUsecase := userUC.NewGetUserAllUsecase(userRepo)
	getStatusUsecase := statusUC.NewGetStatusUsecase()
	pingUsecase := pingUC.NewPingUsecase()
	welcomeUsecase := welcomeUC.NewWelcomeUsecase()

	// 4️⃣ CREAR HANDLERS
	userHandler := handler.NewUserHandler(getUserUsecase, getUserAllUsecase)
	statusHandler := handler.NewStatusHandler(getStatusUsecase)
	pingHandler := handler.NewPingHandler(pingUsecase)
	welcomeHandler := handler.NewWelcomeHandler(welcomeUsecase)
	userNameHanlder := handler.NewUserNameHandler(getUserUsecase)

	// 5️⃣ CONFIGURAR ROUTER
	r := router.Setup(userHandler, statusHandler, pingHandler, welcomeHandler, userNameHanlder)

	// 6️⃣ INICIAR SERVIDOR
	log.Printf("🚀 Servidor iniciado en http://localhost:%s", cfg.ServerPort)
	if err := router.Start(cfg, r); err != nil {
		log.Fatal("❌ Error al iniciar el servidor:", err)
	}
}

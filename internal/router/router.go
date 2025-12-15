package router

import (
	"ejercicio-api/internal/config"
	"ejercicio-api/internal/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Setup configura el router HTTP con los handlers ya inicializados
// Responsabilidad: Configurar rutas, middleware y conectar handlers
func Setup(
	userHandler *handler.UserHandler,
	statusHandler *handler.StatusHandler,
	pingHandler *handler.PingHandler,
	welcomeHandler *handler.WelcomeHandler,
) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Log de cada request
	r.Use(middleware.Recoverer) // Recuperación de panics
	r.Use(middleware.RequestID) // ID único por request

	// Health checks
	r.Get("/status", statusHandler.Get)
	r.Get("/ping", pingHandler.Ping)
	r.Get("/", welcomeHandler.Welcome)

	// Rutas de usuarios
	r.Get("/users", userHandler.GetAll)       // Listar usuarios
	r.Get("/users/{id}", userHandler.GetByID) // Obtener usuario por ID

	return r
}

// Start inicia el servidor HTTP en el puerto configurado
func Start(cfg *config.Config, router *chi.Mux) error {
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	return http.ListenAndServe(addr, router)
}


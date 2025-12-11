package http

import (
	"ejercicio-api/internal/adapter/http/handler"
	"ejercicio-api/internal/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter configura el router HTTP con los handlers ya inicializados
// Esta función pertenece a la capa de INFRAESTRUCTURA
// Responsabilidad: Configurar rutas, middleware y conectar handlers
func SetupRouter(
	userHandler *handler.UserHandler,
	statusHandler *handler.StatusHandler,
	pingHandler *handler.PingHandler,
) *chi.Mux {
	// Crear router chi
	r := chi.NewRouter()

	// Middleware: funciones que se ejecutan antes de los handlers
	r.Use(middleware.Logger)    // Log de cada request
	r.Use(middleware.Recoverer) // Recuperación de panics
	r.Use(middleware.RequestID) // ID único por request

	// REGISTRAR RUTAS
	// Cada ruta conecta un endpoint HTTP con un handler

	// Health checks
	r.Get("/status", statusHandler.Get) // Estado del servidor
	r.Get("/ping", pingHandler.Ping)    // Verificación rápida

	// Rutas de negocio
	r.Get("/users/{id}", userHandler.GetByID) // Obtener usuario por ID

	return r
}

// Start inicia el servidor HTTP en el puerto configurado
func Start(cfg *config.Config, router *chi.Mux) error {
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	return http.ListenAndServe(addr, router)
}

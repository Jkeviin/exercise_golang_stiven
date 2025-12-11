package http

import (
	"ejercicio-api/internal/adapter/http/handler"
	"ejercicio-api/internal/adapter/repository"
	"ejercicio-api/internal/config"
	pingUsecase "ejercicio-api/internal/usecase/ping"
	statusUsecase "ejercicio-api/internal/usecase/status"
	userUsecase "ejercicio-api/internal/usecase/user"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// Status
	statusUC := statusUsecase.NewGetStatusUsecase()
	statusHandler := handler.NewStatusHandler(statusUC)
	r.Get("/status", statusHandler.Get)

	// Ping
	pingUC := pingUsecase.NewPingUsecase()
	pingHandler := handler.NewPingHandler(pingUC)
	r.Get("/ping", pingHandler.Ping)

	// Users
	userRepo := repository.NewUserAPIRepository(cfg.ExternalAPIURL)
	getUserUC := userUsecase.NewGetUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(getUserUC)
	r.Get("/users/{id}", userHandler.GetByID)

	return r
}

func Start(cfg *config.Config, router *chi.Mux) error {
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("🚀 Servidor iniciado en http://localhost%s", addr)
	return http.ListenAndServe(addr, router)
}

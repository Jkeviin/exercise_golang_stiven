package handler

import (
	welcomeUC "ejercicio-api/internal/usecase/welcome"
	"encoding/json"
	"net/http"
)

// WelcomeHandler maneja el endpoint de bienvenida con información de la API
// Propósito: Proporcionar documentación básica e información sobre endpoints disponibles
type WelcomeHandler struct {
	welcomeUC *welcomeUC.WelcomeUsecase
}

func NewWelcomeHandler(welcomeUC *welcomeUC.WelcomeUsecase) *WelcomeHandler {
	return &WelcomeHandler{
		welcomeUC: welcomeUC,
	}
}

// Welcome devuelve información de bienvenida, versión y lista de endpoints
func (h *WelcomeHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	welcome := h.welcomeUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)
}


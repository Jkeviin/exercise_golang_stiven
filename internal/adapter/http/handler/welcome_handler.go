package handler

import (
	welcomeUsecase "ejercicio-api/internal/usecase/welcome"
	"encoding/json"
	"net/http"
)

type WelcomeHandler struct {
	welcomeUC *welcomeUsecase.WelcomeUsecase
}

func NewWelcomeHandler(welcomeUC *welcomeUsecase.WelcomeUsecase) *WelcomeHandler {
	return &WelcomeHandler{
		welcomeUC: welcomeUC,
	}
}

func (h *WelcomeHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	welcome := h.welcomeUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)
}

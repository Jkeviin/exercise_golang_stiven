package handler

import (
	userUsecase "ejercicio-api/internal/usecase/user"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	getUserUC *userUsecase.GetUserUsecase
}

func NewUserHandler(getUserUC *userUsecase.GetUserUsecase) *UserHandler {
	return &UserHandler{
		getUserUC: getUserUC,
	}
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	user, err := h.getUserUC.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

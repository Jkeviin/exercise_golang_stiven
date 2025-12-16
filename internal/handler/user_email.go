package handler

import (
	domainUser "ejercicio-api/internal/domain/user"
	userUC "ejercicio-api/internal/usecase/user"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserEmailHandler struct {
	getUserUC *userUC.GetUserUsecase
}

func NewUserEmailHandler(getUserUC *userUC.GetUserUsecase) *UserEmailHandler {
	return &UserEmailHandler{
		getUserUC: getUserUC,
	}
}

func (h *UserEmailHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, `{"error":"ID inválido"}`, http.StatusBadRequest)
		return
	}

	user, err := h.getUserUC.Execute(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}
	email := domainUser.UserEmail{Email: user.Email}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(email)
}

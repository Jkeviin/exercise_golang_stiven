package handler

import (
	domainUser "ejercicio-api/internal/domain/user"
	userUC "ejercicio-api/internal/usecase/user"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserNameHandler struct {
	getUserUC *userUC.GetUserUsecase
}

func NewUserNameHandler(getUserUC *userUC.GetUserUsecase) *UserNameHandler {
	return &UserNameHandler{
		getUserUC: getUserUC,
	}
}
func (h *UserNameHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id") // aca se trae el id del usuario

	id, err := strconv.Atoi(idParam) // conversion de letras a numeros
	if err != nil {
		http.Error(w, `{"error":"ID inválido"}`, http.StatusBadRequest)
		return
	}

	user, err := h.getUserUC.Execute(id) // Usuario
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	name := domainUser.UserName{Name: user.Name}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(name)
}

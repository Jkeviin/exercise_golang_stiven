package handler

import (
	statusUsecase "ejercicio-api/internal/usecase/status"
	"encoding/json"
	"net/http"
)

type StatusHandler struct {
	getStatusUC *statusUsecase.GetStatusUsecase
}

func NewStatusHandler(getStatusUC *statusUsecase.GetStatusUsecase) *StatusHandler {
	return &StatusHandler{
		getStatusUC: getStatusUC,
	}
}

func (h *StatusHandler) Get(w http.ResponseWriter, r *http.Request) {
	status := h.getStatusUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

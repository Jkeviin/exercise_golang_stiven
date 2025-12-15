package handler

import (
	statusUC "ejercicio-api/internal/usecase/status"
	"encoding/json"
	"net/http"
)

// StatusHandler maneja las peticiones de estado detallado del servidor
// Propósito: Monitoreo con métricas (uptime, versión, ambiente, contador de peticiones)
type StatusHandler struct {
	getStatusUC *statusUC.GetStatusUsecase
}

func NewStatusHandler(getStatusUC *statusUC.GetStatusUsecase) *StatusHandler {
	return &StatusHandler{
		getStatusUC: getStatusUC,
	}
}

// Get devuelve el estado detallado del servidor con métricas
func (h *StatusHandler) Get(w http.ResponseWriter, r *http.Request) {
	status := h.getStatusUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}


package handler

import (
	pingUC "ejercicio-api/internal/usecase/ping"
	"encoding/json"
	"net/http"
)

// PingHandler maneja las peticiones de health check simple
// Propósito: Verificación rápida de que el servidor está respondiendo (usado por monitoreo/load balancers)
type PingHandler struct {
	pingUC *pingUC.PingUsecase
}

func NewPingHandler(pingUC *pingUC.PingUsecase) *PingHandler {
	return &PingHandler{
		pingUC: pingUC,
	}
}

// Ping responde con un mensaje simple para verificar que la API está activa
func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ping := h.pingUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ping)
}


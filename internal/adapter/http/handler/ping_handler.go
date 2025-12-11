package handler

import (
	pingUsecase "ejercicio-api/internal/usecase/ping"
	"encoding/json"
	"net/http"
)

type PingHandler struct {
	pingUC *pingUsecase.PingUsecase
}

func NewPingHandler(pingUC *pingUsecase.PingUsecase) *PingHandler {
	return &PingHandler{
		pingUC: pingUC,
	}
}

func (h *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ping := h.pingUC.Execute()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ping)
}


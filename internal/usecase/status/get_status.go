package status

import (
	"ejercicio-api/internal/domain/status"
	"time"
)

type GetStatusUsecase struct {
	startTime time.Time
}

func NewGetStatusUsecase() *GetStatusUsecase {
	return &GetStatusUsecase{
		startTime: time.Now(),
	}
}

var requesCount int

func (uc *GetStatusUsecase) Execute() *status.Status {
	uptime := time.Since(uc.startTime).Seconds()
	requesCount++
	return &status.Status{
		Message:       "La aplicación está funcionando correctamente",
		Version:       "1.1.0",
		Uptime:        int64(uptime),
		Environment:   "development",
		Timestamp:     time.Now().Format(time.RFC3339),
		RequestAcount: requesCount,
	}
}

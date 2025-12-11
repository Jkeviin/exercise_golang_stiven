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

func (uc *GetStatusUsecase) Execute() *status.Status {
	uptime := time.Since(uc.startTime).Seconds()

	return &status.Status{
		Message: "La aplicación está funcionando correctamente",
		Version: "1.0.0",
		Uptime:  int64(uptime),
	}
}

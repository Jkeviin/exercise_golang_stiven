package ping

import "ejercicio-api/internal/domain/ping"

type PingUsecase struct{}

func NewPingUsecase() *PingUsecase {
	return &PingUsecase{}
}

func (uc *PingUsecase) Execute() *ping.Ping {
	return &ping.Ping{
		Message: "API funcionando correctamente",
	}
}

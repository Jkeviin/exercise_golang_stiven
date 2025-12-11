package ping_test

import (
	pingUsecase "ejercicio-api/internal/usecase/ping"
	"testing"
)

func TestPingUsecase_Execute(t *testing.T) {
	uc := pingUsecase.NewPingUsecase()
	ping := uc.Execute()

	if ping.Message != "pong" {
		t.Errorf("Se esperaba 'pong', pero se obtuvo '%s'", ping.Message)
	}
}


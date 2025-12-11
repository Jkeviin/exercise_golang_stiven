package status_test

import (
	statusUsecase "ejercicio-api/internal/usecase/status"
	"testing"
)

func TestGetStatusUsecase_Execute(t *testing.T) {
	uc := statusUsecase.NewGetStatusUsecase()
	status := uc.Execute()

	if status.Message == "" {
		t.Error("El mensaje no debería estar vacío")
	}

	if status.Version != "1.0.0" {
		t.Errorf("Se esperaba versión 1.0.0, pero se obtuvo %s", status.Version)
	}

	if status.Uptime < 0 {
		t.Error("Uptime no puede ser negativo")
	}
}


package user_test

import (
	"ejercicio-api/internal/domain/user"
	userUsecase "ejercicio-api/internal/usecase/user"
	"testing"
)

type MockUserRepository struct{}

func (m *MockUserRepository) FindByID(id int) (*user.User, error) {
	return &user.User{
		ID:       id,
		Name:     "Usuario de Prueba",
		Email:    "test@example.com",
		Username: "testuser",
	}, nil
}

func TestGetUserUsecase_Execute(t *testing.T) {
	repo := &MockUserRepository{}
	uc := userUsecase.NewGetUserUsecase(repo)

	user, err := uc.Execute(1)
	if err != nil {
		t.Fatalf("No se esperaba error: %v", err)
	}

	if user.Name == "" {
		t.Error("El nombre no debería estar vacío")
	}
}

func TestGetUserUsecase_Execute_InvalidID(t *testing.T) {
	repo := &MockUserRepository{}
	uc := userUsecase.NewGetUserUsecase(repo)

	_, err := uc.Execute(0)
	if err == nil {
		t.Error("Se esperaba un error para ID = 0")
	}
}

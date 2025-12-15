package repository

import (
	"ejercicio-api/internal/domain/user"
)

// UserRepository define el contrato para operaciones con usuarios
type UserRepository interface {
	FindByID(id int) (*user.User, error)
	FindAll() ([]*user.User, error)
}


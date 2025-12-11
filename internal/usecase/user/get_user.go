package user

import (
	"ejercicio-api/internal/domain/user"
	"fmt"
)

type GetUserUsecase struct {
	userRepo user.Repository
}

func NewGetUserUsecase(userRepo user.Repository) *GetUserUsecase {
	return &GetUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *GetUserUsecase) Execute(id int) (*user.User, error) {
	if id <= 0 {
		return nil, fmt.Errorf("el ID debe ser mayor que 0")
	}

	return uc.userRepo.FindByID(id)
}

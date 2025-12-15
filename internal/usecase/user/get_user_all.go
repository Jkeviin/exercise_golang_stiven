package user

import (
	"ejercicio-api/internal/domain/user"
)

type GetUserAllUsecase struct {
	userRepo user.Repository
}

func NewGetUserAllUsecase(userRepo user.Repository) *GetUserAllUsecase {
	return &GetUserAllUsecase{
		userRepo: userRepo,
	}
}

func (uc *GetUserAllUsecase) Execute() (*user.Users, error) {

	return uc.userRepo.FindAll()
}

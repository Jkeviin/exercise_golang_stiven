package user

import (
	"ejercicio-api/internal/domain/user"
	"ejercicio-api/internal/repository"
)

type GetUserAllUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserAllUsecase(userRepo repository.UserRepository) *GetUserAllUsecase {
	return &GetUserAllUsecase{
		userRepo: userRepo,
	}
}

func (uc *GetUserAllUsecase) Execute() ([]*user.User, error) {
	return uc.userRepo.FindAll()
}

package user

import (
	"ejercicio-api/internal/domain/user"
	"ejercicio-api/internal/repository"
	"fmt"
)

type GetUserUsecase struct {
	userRepo repository.UserRepository
}

func NewGetUserUsecase(userRepo repository.UserRepository) *GetUserUsecase {
	return &GetUserUsecase{
		userRepo: userRepo,
	}
}

func (uc *GetUserUsecase) Execute(id int) (*user.User, error) {
	if id <= 0 { // validacion menor que 0 pasa tan
		return nil, fmt.Errorf("el ID debe ser mayor que 0")
	} else if id > 10 { // validacion si es mayor que 10 pasa tan
		return nil, fmt.Errorf("el ID debe ser menor que 10")
	}
	return uc.userRepo.FindByID(id) //Usuario
}

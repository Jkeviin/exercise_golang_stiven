package welcome

import "ejercicio-api/internal/domain/welcome"

type WelcomeUsecase struct{}

func NewWelcomeUseCase() *WelcomeUsecase {
	return &WelcomeUsecase{}
}

func (uc *WelcomeUsecase) Execute() *welcome.Welcome {
	return &welcome.Welcome{
		Message: "Bienvenido a la API de Ejercicio",
		Version: 1.0,
		Endpoints: []string{
			"/status",
			"/ping",
			"/user",
			"/users/{id}",
		},
	}
}

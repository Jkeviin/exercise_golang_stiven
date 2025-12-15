package repository

import (
	"ejercicio-api/internal/domain/user"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserAPIRepository struct {
	baseURL string
	client  *http.Client
}

// NewUserRepository crea una nueva instancia del repositorio de usuarios
func NewUserRepository(baseURL string) UserRepository {
	return &UserAPIRepository{
		baseURL: baseURL,
		client:  &http.Client{},
	}
}

func (r *UserAPIRepository) FindByID(id int) (*user.User, error) {
	url := fmt.Sprintf("%s/users/%d", r.baseURL, id)

	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al consultar API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("el servidor externo no está disponible")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API respondió con código: %d", resp.StatusCode)
	}

	var u user.User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("error al decodificar respuesta: %w", err)
	}

	return &u, nil
}

func (r *UserAPIRepository) FindAll() ([]*user.User, error) {
	url := fmt.Sprintf("%s/users", r.baseURL)
	
	resp, err := r.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error al consultar API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("el servidor externo no está disponible")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API respondió con código: %d", resp.StatusCode)
	}

	var users []*user.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, fmt.Errorf("error al decodificar respuesta: %w", err)
	}

	return users, nil
}


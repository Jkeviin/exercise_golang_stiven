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

func NewUserAPIRepository(baseURL string) *UserAPIRepository {
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

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API respondió con código: %d", resp.StatusCode)
	}

	var u user.User
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("error al decodificar respuesta: %w", err)
	}

	return &u, nil
}

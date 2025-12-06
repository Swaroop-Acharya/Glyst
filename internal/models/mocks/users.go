package mocks

import (
	"time"

	"glyst/internal/models"
)

type UserModel struct{}

func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "swaroop@gmail.com" && password == "password" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (m *UserModel) Get(id int) (models.User, error) {
	switch id {
	case 1:
		u := models.User{
			ID:      1,
			Name:    "Swaroop",
			Email:   "swaroop@test.com",
			Created: time.Now(),
		}
		return u, nil
	default:
		return models.User{}, models.ErrNoRecord
	}
}

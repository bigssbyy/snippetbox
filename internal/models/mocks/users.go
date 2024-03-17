package mocks

import (
	"time"

	"github.com/bigssbyy/snippetbox/internal/models"
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
	if email == "alice@example.com" && password == "pa$$word" {
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

	var user models.User

	user.ID = 1
	user.Name = "Carol Smith"
	user.Email = "carol@example.com"
	user.Created = time.Now().UTC()

	return user, nil
}

func (m *UserModel) Update(id int, newPassword string) error {

	var user models.User

	user.ID = 1
	user.Name = "Carol Smith"
	user.Email = "carol@example.com"
	user.Created = time.Now().UTC()

	return nil
}

func (m *UserModel) ValidatePassword(id int, newPassword string) error {

	var user models.User

	user.ID = 1
	user.Name = "Carol Smith"
	user.Email = "carol@example.com"
	user.Created = time.Now().UTC()

	return nil
}

package store

import (
	"github.com/NikolayStrekalov/practicum-gophermart/internal/db"
	"github.com/NikolayStrekalov/practicum-gophermart/internal/models"
)

func RegisterUser(r *models.Registration) (*models.User, error) {
	db.SaveUser(r.Login, r.Password)
	// TODO
	return nil, nil
}

func User(id int) models.User {
	return models.User{}
}

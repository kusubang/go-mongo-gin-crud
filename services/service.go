package services

import "go-mongodb/models"

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetAUser(userId string) (models.User, error)
	CreateAUser(user *models.User) (*models.User, error)
	EditAUser(userId string, user *models.User) (*models.User, error)
	DeleteAUser(userId string) error
}

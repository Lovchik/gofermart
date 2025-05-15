package storage

import "gofermart/internal/server/models"

type Storage interface {
	HealthCheck() error
	IsUserExists(request models.LoginRequest) bool
	GetUserByCreds(request models.LoginRequest) models.User
	CreateUser(request models.LoginRequest) (models.User, error)
}

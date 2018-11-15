package user

import (
	"context"
	"github.com/bottlenome/ll3/models"
)

type UserRepository interface {
	GetByUserName(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

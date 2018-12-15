package user

import (
	"github.com/bottlenome/ll3/models"
)

type UserRepository interface {
	GetByUserName(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	CalcTotalMony() (total float64, err error)
}

package managers

import "github.com/anshidmattara7861/Go-Gin-backend/models"

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(user *models.User) (*models.User, error) {
	return nil, nil
}

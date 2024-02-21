package managers

import (
	"errors"

	"github.com/anshidmattara7861/Go-Gin-backend/common"
	"github.com/anshidmattara7861/Go-Gin-backend/database"
	"github.com/anshidmattara7861/Go-Gin-backend/models"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (userMgr *UserManager) Create(userData *common.UserCreationInput) (*models.User, error) {

	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newUser, nil
}

func (userMgr *UserManager) List() ([]models.User, error) {

	users := []models.User{}

	database.DB.Find(&users)

	return users, nil
}

func (userMgr *UserManager) Details(userId string) (models.User, error) {

	user := models.User{}

	database.DB.First(&user, userId)

	return user, nil
}

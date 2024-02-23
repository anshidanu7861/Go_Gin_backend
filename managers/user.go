package managers

import (
	"errors"

	"github.com/anshidmattara7861/Go-Gin-backend/common"
	"github.com/anshidmattara7861/Go-Gin-backend/database"
	"github.com/anshidmattara7861/Go-Gin-backend/models"
)

type UserManager interface {
	Create(userData *common.UserCreationInput) (*models.User, error)
	List() ([]models.User, error)
	Details(userId string) (models.User, error)
	Update(userId string, userData *common.UserUpdateInput) (*models.User, error)
	Delete(userId string) ( error) 
}

type userManager struct {
}

func NewUserManager() UserManager {
	return &userManager{}
}

func (userMgr *userManager) Create(userData *common.UserCreationInput) (*models.User, error) {

	newUser := &models.User{FullName: userData.FullName, Email: userData.Email}
	database.DB.Create(newUser)

	if newUser.ID == 0 {
		return nil, errors.New("user creation failed")
	}

	return newUser, nil
}


func (userMgr *userManager) List() ([]models.User, error) {

	users := []models.User{}

	database.DB.Find(&users)

	return users, nil
}

func (userMgr *userManager) Details(userId string) (models.User, error) {

	user := models.User{}

	database.DB.First(&user, userId)

	return user, nil
}

func (userMgr *userManager) Update(userId string, userData *common.UserUpdateInput) (*models.User, error) {


	user :=  models.User{}

	database.DB.First(&user, userId)

	// user.FullName = userData.FullName
	// user.Email = userData.Email

	// database.DB.Save(&user)

	database.DB.Model(&user).Updates(models.User{FullName: userData.FullName, Email: userData.Email})

	return &user, nil
}


func (userMgr *userManager) Delete(userId string) ( error) {

	user := models.User{}

	database.DB.First(&user, userId)

	database.DB.Delete(&user)

	return nil
}
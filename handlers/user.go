package handlers

import (
	"fmt"
	"net/http"

	"github.com/anshidmattara7861/Go-Gin-backend/common"
	"github.com/anshidmattara7861/Go-Gin-backend/managers"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	groupName   string
	UserManager *managers.UserManager
}

func NewUserHandlerFrom(userManager *managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users",
		userManager,
	}
}

func (userHandler *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(userHandler.groupName)
	userGroup.POST("", userHandler.Create)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {

	userData := common.NewUserCreationInput()

	err := ctx.BindJSON(&userData)

	if err != nil {
		fmt.Println("filed to bind data")
	}

	newUser, err := userHandler.UserManager.Create(userData)

	if err != nil {
		fmt.Println("failed to creations")
	}

	ctx.JSON(http.StatusOK, newUser)
}

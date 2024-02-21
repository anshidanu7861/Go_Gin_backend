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
	userGroup.GET("", userHandler.List)
	userGroup.GET(":userId/", userHandler.Details)
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

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := userHandler.UserManager.List()
	if err != nil {
		fmt.Println("failed to find users")
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (userHandler *UserHandler) Details(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		fmt.Println("Invalid userId")
	}

	userDetails, err := userHandler.UserManager.Details(userId)
	if err != nil {
		fmt.Println("failed to find users")
	}
	ctx.JSON(http.StatusOK, userDetails)
}

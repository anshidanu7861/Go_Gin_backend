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
	userGroup.DELETE(":userId/", userHandler.Delete)
}

func (userHandler *UserHandler) Create(ctx *gin.Context) {
	userData := common.NewUserCreationInput()
	err := ctx.BindJSON(&userData)
	if err != nil {
		common.FailMessage(ctx, "failed to bind data")
		return
	}
	newUser, err := userHandler.UserManager.Create(userData)
	if err != nil {
		common.FailMessage(ctx, "failed to creations")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

func (userHandler *UserHandler) List(ctx *gin.Context) {

	allUsers, err := userHandler.UserManager.List()
	if err != nil {
		common.FailMessage(ctx, "failed to find users")
		return 
	}
	ctx.JSON(http.StatusOK, allUsers)
}

func (userHandler *UserHandler) Details(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		common.FailMessage(ctx, "Invalid user id")
		return 
	}

	userDetails, err := userHandler.UserManager.Details(userId)

	if userDetails.ID == 0 {
		common.FailMessage(ctx, "user not found")
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}

func (userHandler *UserHandler) Delete(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		common.FailMessage(ctx, "invalid user id")
		return
	}

	err := userHandler.UserManager.Delete(userId)
	if err != nil {
		fmt.Println("failed to find users")
	}
	
	
	common.SuccessMessage(ctx, "Deleted successfully")
}

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
	UserManager managers.UserManager
}

func NewUserHandlerFrom(userManager managers.UserManager) *UserHandler {
	return &UserHandler{
		"api/users",
		userManager,
	}
}

func (handlers *UserHandler) RegisterUserApis(r *gin.Engine) {
	userGroup := r.Group(handlers.groupName)
	userGroup.POST("", handlers.Create)
	userGroup.GET("", handlers.List)
	userGroup.GET(":userId/", handlers.Details)
	userGroup.DELETE(":userId/", handlers.Delete)
	userGroup.PATCH(":userId/", handlers.Update)
}


// USER CREATION
func (handlers *UserHandler) Create(ctx *gin.Context) {
	userData := common.NewUserCreationInput()
	err := ctx.BindJSON(&userData)
	if err != nil {
		common.FailMessage(ctx, "failed to bind data")
		return
	}
	newUser, err := handlers.UserManager.Create(userData)
	if err != nil {
		common.FailMessage(ctx, "failed to creations")
		return
	}
	ctx.JSON(http.StatusOK, newUser)
}

// ALL USERS LIST
func (handlers *UserHandler) List(ctx *gin.Context) {

	allUsers, err := handlers.UserManager.List()
	if err != nil {
		common.FailMessage(ctx, "failed to find users")
		return 
	}
	ctx.JSON(http.StatusOK, allUsers)
}


// USER SINGLE DETAILS
func (handlers *UserHandler) Details(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		common.FailMessage(ctx, "Invalid user id")
		return 
	}

	userDetails, err := handlers.UserManager.Details(userId)

	if err != nil {
		fmt.Println(err, "get error")
	}

	if userDetails.ID == 0 {
		common.FailMessage(ctx, "user not found")
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}


// DELETE USER
func (handlers *UserHandler) Delete(ctx *gin.Context) {

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		common.FailMessage(ctx, "invalid user id")
		return
	}

	err := handlers.UserManager.Delete(userId)
	if err != nil {
		fmt.Println("failed to find users")
	}
	
	
	common.SuccessMessage(ctx, "Deleted successfully")
}

// USER UPDATE
func (handlers *UserHandler) Update(ctx *gin.Context) {
	userData := common.NewUserUpdateInput()
	err := ctx.BindJSON(&userData)
	if err != nil {
		common.FailMessage(ctx, "failed to bind data")
		return
	}

	userId, ok := ctx.Params.Get("userId")

	if !ok {
		common.FailMessage(ctx, "invalid user id")
		return
	}
	updateUser, err := handlers.UserManager.Update(userId,userData)


	if err != nil {
		common.FailMessage(ctx, "failed to update user")
		return
	}
	
	ctx.JSON(http.StatusOK, updateUser)
}


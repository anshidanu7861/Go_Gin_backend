package handlers

import (
	"net/http"

	"github.com/anshidmattara7861/Go-Gin-backend/database"
	"github.com/anshidmattara7861/Go-Gin-backend/managers"
	"github.com/anshidmattara7861/Go-Gin-backend/models"
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

	database.DB.Create(&models.User{FullName: "Tom", Email: "example@gmail.com"})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "api version",
	})
}

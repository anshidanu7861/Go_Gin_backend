package handlers

import (
	"fmt"
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

	var userData struct {
		fullName string `json:"full_name"`
		email    string `json: "email"`
	}

	err := ctx.BindJSON(userData)

	fmt.Println(userData.email)
	fmt.Println(userData.fullName)

	if err != nil {
		fmt.Println("filed to bind data")
	}

	database.DB.Create(&models.User{FullName: "Tom", Email: "example@gmail.com"})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "api version",
	})
}

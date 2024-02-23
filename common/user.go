package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCreationInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type UserUpdateInput struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func NewUserCreationInput() *UserCreationInput {
	return &UserCreationInput{}
}

func NewUserUpdateInput() *UserUpdateInput {
	return &UserUpdateInput{}
}


type requestResponse struct {
	Message string `json:"message"`
	Status uint `json:"status"`
}

func SuccessMessage(ctx *gin.Context,msg string){

	response := requestResponse{
		Message: msg,
		Status: http.StatusOK,
	}

	ctx.JSON(http.StatusOK, response)
}

func FailMessage(ctx *gin.Context,msg string){

	response := requestResponse{
		Message: msg,
		Status: http.StatusBadGateway,
	}

	ctx.JSON(http.StatusBadGateway, response)
}
package main

import (
	"github.com/anshidmattara7861/Go-Gin-backend/database"
	"github.com/anshidmattara7861/Go-Gin-backend/handlers"
	"github.com/anshidmattara7861/Go-Gin-backend/managers"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Initialize()

	router := gin.Default()

	userManger := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterUserApis(router)

	router.Run()
}

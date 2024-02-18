package main

import (
	"fmt"

	"github.com/anshidmattara7861/Go-Gin-backend/database"
	"github.com/anshidmattara7861/Go-Gin-backend/handlers"
	"github.com/anshidmattara7861/Go-Gin-backend/managers"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println(database.DB, "show this ")
	database.Initialize()
}

func main() {

	router := gin.Default()

	userManger := managers.NewUserManager()
	userHandler := handlers.NewUserHandlerFrom(userManger)
	userHandler.RegisterUserApis(router)

	router.Run()
}

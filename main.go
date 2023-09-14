package main

import (
	"inventory-api/controller"
	"inventory-api/initiliazer"
	"inventory-api/repository"
	"inventory-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	db, err = initiliazer.ConnectToDatabase()
	err = initiliazer.SyncDatabase(db)
	if err != nil {
		log.Fatal("Connection Database Failed")
	}
}

func main() {

	//adding Repository
	userRepository := repository.NewUserRepository(db)

	//adding Service
	UserService := services.NewUserService(userRepository)

	//adding Controller
	userController := controller.NewUserController(UserService)

	//Routing
	router := gin.Default()

	//Routing Grouping
	// routerLogin := router.Group("/v1")

	router.POST("/signup", userController.SignUp)

	router.Run(":8080")
}

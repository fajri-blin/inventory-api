package main

import (
	"inventory-api/controller"
	"inventory-api/initiliazer"
	"inventory-api/repository"
	"inventory-api/services"
	"inventory-api/middleware"
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
	trxRepo := repository.NewTransactionRepository(db)

	//adding Service
	UserService := services.NewUserService(userRepository)
	trxService := services.NewTransactionService(trxRepo)

	//adding Controller
	userController := controller.NewUserController(UserService)
	trxController := controller.NewTransactionController(trxService)

	//Routing
	router := gin.Default()

	//Routing Grouping
	routerLogin := router.Group("/user", middleware.RequireAuth)

	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)

	routerLogin.POST("/create/supplier")

	router.POST("/transaction", trxController.PostTrxController)

	router.Run(":8080")
}

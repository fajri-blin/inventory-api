package main

import (
	"inventory-api/controller"
	"inventory-api/initiliazer"
	"inventory-api/middleware"
	"inventory-api/repository"
	"inventory-api/services"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error

	initiliazer.LoadEnvVariables()
	db, err = initiliazer.ConnectToDatabase()
	err = initiliazer.SyncDatabase(db)
	if err != nil {
		log.Fatal("Connection Database Failed")
	}
}

func main() {

	//adding Repository
	userRepository := repository.NewUserRepository(db)
	supplierRepository := repository.NewSupplierRepository(db)
	productRepository := repository.NewProductRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	//adding Service
	UserService := services.NewUserService(userRepository)
	supplierService := services.NewSupplierService(supplierRepository)
	productService := services.NewProductService(productRepository)
	transactionService := services.NewTransactionService(transactionRepository)

	//adding Controller
	userController := controller.NewUserController(UserService)
	supplierController := controller.NewSupplierController(supplierService)
	productController := controller.NewProductController(productService)
	transactionController := controller.NewTransactionController(transactionService)

	//Routing
	router := gin.Default()

	// Grouping User
	routerUser := router.Group("/user", middleware.RequireAuth)
	routerUser.POST("/signup", userController.SignUp)
	routerUser.POST("/login", userController.Login)

	// Grouping Product
	routerProduct := router.Group("/product", middleware.RequireAuth)
	routerProduct.POST("/create", productController.Create)
	routerProduct.PUT("/update/:id", productController.Update)
	routerProduct.DELETE("/delete/:id", productController.Delete)
	routerProduct.GET("/", productController.GetAll)
	routerProduct.GET("/:id", productController.GetByID)

	// Grouping Supplier
	routerSupplier := router.Group("/supplier", middleware.RequireAuth)
	routerSupplier.POST("/create", supplierController.CreateCompanyController)
	/* 	routerSupplier.PUT("/update/:id", supplierController.Update)
	   	routerSupplier.DELETE("/delete/:id", supplierController.Delete)
	   	routerSupplier.GET("/", supplierController.GetAll)
	   	routerSupplier.GET("/:id", supplierController.GetByID) */

	// Grouping Transaction
	routerTrx := router.Group("/transaction", middleware.RequireAuth)
	routerTrx.POST("/create", transactionController.Create)
	routerTrx.PUT("/update/:id", transactionController.Update)
	routerTrx.DELETE("/delete/:id", transactionController.Delete)
	routerTrx.GET("/", transactionController.GetAll)
	routerTrx.GET("/:id", transactionController.GetByID)

	// Run
	router.Run(":8080")
}

package main

import (
	"fmt"
	"log"

	"dbo/controllers"
	"dbo/database"
	"dbo/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	// database.Database1.AutoMigrate(&entity.Customer{})
	database.ConnectMYSQL()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controllers.Register)
	publicRoutes.POST("/login", controllers.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/customers", controllers.FindCustomers)
	protectedRoutes.GET("/customers/:id", controllers.FindCustomerById)
	protectedRoutes.POST("/customers", controllers.CreateCustomer)
	protectedRoutes.PUT("/customers", controllers.UpdateCustomer)
	protectedRoutes.DELETE("/customers/:id", controllers.DeleteCustomer)

	protectedRoutes.GET("/orders", controllers.FindOrders)
	protectedRoutes.GET("/orders/:id", controllers.FindOrderById)
	protectedRoutes.POST("/orders", controllers.CreateOrder)
	// protectedRoutes.PUT("/customers", controllers.UpdateCustomer)
	// protectedRoutes.DELETE("/customers/:id", controllers.DeleteCustomer)

	router.Run("localhost:8000")
	fmt.Println("Server running on port 8000")
}

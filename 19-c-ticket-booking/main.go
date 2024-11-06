package main

import (
	"fmt"
	"ticket-booking/controllers"
	"ticket-booking/database"
	"ticket-booking/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	database.InitDB()

	// Initialize Stripe
	controllers.InitStripe()

	// Initialize Gin
	r := gin.Default()

	// Public routes
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.GET("/events", controllers.ListEvents)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/book", controllers.BookTicket)
		protected.POST("/pay", controllers.ProcessPayment) // New payment route
	}

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	r.Run(":8080")
}

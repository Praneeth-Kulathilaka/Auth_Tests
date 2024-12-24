package main

import (
	// "my-gin-backend/routes"
	// "my-gin-backend/services"
	"github.com/gin-contrib/cors"
	"github.com/Praneeth-Kulathilaka/Auth_Tests/routes"
	"github.com/Praneeth-Kulathilaka/Auth_Tests/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Firebase
	services.InitFirebase()

	// Set up Gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8080"},
        AllowMethods:     []string{"GET", "POST"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

	// Set the HTML template
	r.LoadHTMLFiles("static/index.ejs")

	// Set routes
	routes.AuthRoutes(r)

	// Start server
	r.Run(":8080")
}

package routes

import (
	"github.com/Praneeth-Kulathilaka/Auth_Tests/controllers"
	"github.com/gin-gonic/gin"
	// "test/controllers"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/", controllers.RenderLoginPage)
	router.POST("/verify-token", controllers.VerifyGoogleToken)
}

package controllers

import (
	"context"
	"net/http"

	// "firebase.google.com/go/auth"
	"github.com/Praneeth-Kulathilaka/Auth_Tests/services"
	"github.com/gin-gonic/gin"
	// "test/services"
)

func RenderLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.ejs", gin.H{})
}

func VerifyGoogleToken(c *gin.Context) {
	var body struct {
		IDToken string `json:"idToken"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	client, err := services.FirebaseApp.Auth(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Firebase Auth client error"})
		return
	}

	token, err := client.VerifyIDToken(context.Background(), body.IDToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token verified successfully",
		"uid":     token.UID,
	})
}

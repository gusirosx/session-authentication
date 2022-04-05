package handlers

import (
	"fmt"
	"log"
	"net/http"
	"session-authentication/entity"
	"session-authentication/models"

	"github.com/gin-gonic/gin"
)

// HomePage handler function
func Home(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie("session")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, claims, err := models.GetTokenWithClaims(cookie.Value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//Return the welcome message to the user, along with their username given in the token
	ctx.JSON(http.StatusOK, gin.H{"response": "Hello " + claims.UserName})

}

// Refresh handler function
func Refresh(ctx *gin.Context) {

	// Get user function will be here
	var user *entity.User = &entity.UserDEMO

	token, err := models.GetToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Set the new token in Authorization header
	ctx.Set("Authorization", token)

	if err := CreateCookie(ctx); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "Cookie Refreshed"})
}

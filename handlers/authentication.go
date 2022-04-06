package handlers

import (
	"fmt"
	"log"
	"net/http"
	"session-authentication/entity"
	"session-authentication/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Login handler function
func Login(ctx *gin.Context) {

	var credentials entity.Credentials
	// Get the JSON body and decode into credentials
	if err := ctx.BindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user e-mail is a valid e-mail
	if err := models.ValidateEmail(credentials.Email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check credentials and return the selected user
	user, err := models.Login(credentials)
	if err != nil {
		if err.Error() == "user e-mail is incorrect" || err.Error() == "user password is incorrect" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.Set("Authorization", user.Token)

	if err := CreateCookie(ctx); err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "Helo " + user.FirstName})
}

// Create Cookie function
func CreateCookie(ctx *gin.Context) error {

	token, ok := ctx.Get("Authorization")
	if ok {
		if err := models.ValidateToken(token.(string)); err != nil {
			return err
		} else {
			expiration := time.Now().Add(time.Hour)
			http.SetCookie(ctx.Writer, &http.Cookie{
				Name:     "session",
				Value:    token.(string),
				Path:     "/",
				Expires:  expiration,
				HttpOnly: true,
				//Secure:   true,
			})
			return nil
		}
	} else {
		return fmt.Errorf("no authorization header provided")
	}
}

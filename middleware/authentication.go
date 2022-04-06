package middleware

import (
	"net/http"
	"session-authentication/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Authentication : verify all authorized operations
func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// The session cookie is obtained from the requests ctx, which come with every request
		cookie, err := ctx.Request.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No session cookie header provided"})
				return
			}
			// For any other type of error, return a bad request status
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if !isExpired(cookie.Expires) {
			ctx.SetCookie("session", "", 0, "/", "", true, true) // Delete cookie
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": "Cookie is expired"})
			return
		}
		if cookie.Value != "" {
			err := models.ValidateToken(cookie.Value)
			if err != nil {
				// If the session token is not present, return an unauthorized error
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired or is invalid"})
				return
			} else {
				// Cookie is valid, let the request continue
				return
			}
		} else {
			// If the cookie value is not set, return an unauthorized status
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Cookie value is empty"})
			return
		}
	}
}

// Method to determine if the session has expired
func isExpired(expiry time.Time) bool {
	return expiry.Before(time.Now())
}

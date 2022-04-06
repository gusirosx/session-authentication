package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/* The client need to know that the cookie is expired.
   In the response, we set the session token to an
   empty value and set its expiry as the current time */
func Logout(ctx *gin.Context) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "session",
		Value:   "",
		MaxAge:  0,
		Expires: time.Now(),
	})
	//http.Redirect(ctx.Writer, ctx.Request, "login", http.StatusFound)
}

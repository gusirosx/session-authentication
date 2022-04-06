package routes

import (
	"log"
	"net/http"
	"session-authentication/handlers"
	"session-authentication/middleware"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func RoutesSetup() {
	// Set Gin to production mode
	// gin.SetMode(gin.ReleaseMode)
	// Set up a http server
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Run the http server
	if err := router.Run(port); err != nil {
		log.Fatalln("could not run server: ", err.Error())
	} else {
		log.Println("Server listening on port: ", port)
	}
}

func initializeRoutes(router *gin.Engine) {
	// Handle the index route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})
	// Handle the no route case
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.POST("/login", handlers.Login)
	router.Use(middleware.Authentication())
	router.GET("/logout", handlers.Logout)
	router.GET("/refresh", handlers.Refresh)
	router.GET("/home", handlers.Home)
}

package api

import (
	handlers "main/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes
func SetupRouter() *gin.Engine {
	app := gin.Default()
	app.Static("/img", "./download")
	app.GET("/get", handlers.GetWallpaper)

	return app
}

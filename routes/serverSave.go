package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogihung/go-gin-example/controllers"
)

// InitializeServerSaveRoutes use to initialize routes for managing a bulk update
func InitializeServerSaveRoutes(router *gin.Engine) {
	router.POST("/server/save", controllers.ServerSaveCreate)
}

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogihung/go-gin-example/controllers"
)

// InitializeHostRoutes use to initialize routes for managing a host resource
func InitializeHostRoutes(router *gin.Engine) {
	router.GET("/hosts", controllers.HostIndex)
	router.GET("/hosts/:id", controllers.HostShow)
	router.POST("/hosts", controllers.HostCreate)
	router.PUT("/hosts/:id", controllers.HostUpdate)
	router.PATCH("/hosts/:id", controllers.HostUpdate)
	router.DELETE("/hosts/:id", controllers.HostDelete)
}

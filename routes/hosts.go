package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogihung/gin-example/controllers"
)

// InitializeHostRoutes info coming soon
func InitializeHostRoutes(router *gin.Engine) {
	router.GET("/hosts", controllers.HostIndex)
	router.POST("/hosts", controllers.HostCreate)
	router.PUT("/hosts/:id", controllers.HostUpdate)
	router.PATCH("/hosts/:id", controllers.HostUpdate)
	router.DELETE("/hosts/:id", controllers.HostDelete)
}

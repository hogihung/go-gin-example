package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hogihung/gin-example/controllers"
)

// InitializePodRoutes use to initialize routes for managing a pod resource
func InitializePodRoutes(router *gin.Engine) {
	router.GET("/pods", controllers.PodIndex)
	router.GET("/pods/:id", controllers.PodShow)
	router.POST("/pods", controllers.PodCreate)
	router.PUT("/pods/:id", controllers.PodUpdate)
	router.PATCH("/pods/:id", controllers.PodUpdate)
	router.DELETE("/pods/:id", controllers.PodDelete)
}

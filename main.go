package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	fmt.Println("Spinning up Nimbulus Meta Server.....")
	router = gin.Default()

	initializeRoutes()

	router.Run()
}

func initializeRoutes() {
	router.GET("/hosts", hostIndex)
	router.POST("/hosts", hostCreate)
	router.PUT("/hosts/:id", hostUpdate)
	router.PATCH("/hosts/:id", hostUpdate)
	router.DELETE("/hosts/:id", hostDelete)
}

func hostIndex(c *gin.Context) {
	c.String(http.StatusOK, "called hostIndex")
}

func hostCreate(c *gin.Context) {
	c.String(http.StatusOK, "Called hostCreate")
}

func hostUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Called hostUpdate")
}

func hostDelete(c *gin.Context) {
	c.String(http.StatusOK, "Called hostDelete")
}

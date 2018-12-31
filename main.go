package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hogihung/go-gin-example/routes"
)

var router *gin.Engine

func main() {
	fmt.Println("Spinning up Nimbulus Meta Server.....")
	router = gin.Default()

	routes.InitializeHostRoutes(router)
	routes.InitializePodRoutes(router)

	router.Run()
}

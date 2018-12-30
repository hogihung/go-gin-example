package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller/Handler for Pod based resources
// ...
// blah
// --------------------------------------------------------

// PodIndex used to list all Pods
func PodIndex(c *gin.Context) {
	c.String(http.StatusOK, "called PodIndex")
}

// PodShow used to show a Pod
func PodShow(c *gin.Context) {
	c.String(http.StatusOK, "called PodShow")
}

// PodCreate used to create a new Pod
func PodCreate(c *gin.Context) {
	c.String(http.StatusOK, "Called PodCreate")
}

// PodUpdate used to update an existing Pod
func PodUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Called PodUpdate")
}

// PodDelete used to remove a Pod
func PodDelete(c *gin.Context) {
	c.String(http.StatusOK, "Called PodDelete")
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller/Handler for Host based resources
// ...
// blah
// --------------------------------------------------------

// HostIndex used to list all hosts
func HostIndex(c *gin.Context) {
	c.String(http.StatusOK, "called HostIndex")
}

// HostShow used to show a host
func HostShow(c *gin.Context) {
	c.String(http.StatusOK, "called HostShow")
}

// HostCreate used to create a new host
func HostCreate(c *gin.Context) {
	c.String(http.StatusOK, "Called HostCreate")
}

// HostUpdate used to update an existing host
func HostUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Called HostUpdate")
}

// HostDelete used to remove a host
func HostDelete(c *gin.Context) {
	c.String(http.StatusOK, "Called HxostDelete")
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// description here

func HostIndex(c *gin.Context) {
	c.String(http.StatusOK, "called hostIndex")
}

func HostCreate(c *gin.Context) {
	c.String(http.StatusOK, "Called hostCreate")
}

func HostUpdate(c *gin.Context) {
	c.String(http.StatusOK, "Called hostUpdate")
}

func HostDelete(c *gin.Context) {
	c.String(http.StatusOK, "Called hostDelete")
}

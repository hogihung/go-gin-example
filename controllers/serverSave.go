package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ServerSaveCreate used for bulk processing
func ServerSaveCreate(c *gin.Context) {
	//c.String(http.StatusOK, "Called ServerSaveCreate")
	c.JSON(http.StatusOK, gin.H{"status": "OK", "num_rows": "24"})
}

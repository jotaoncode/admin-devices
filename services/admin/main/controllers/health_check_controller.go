package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// IsHealthy godoc
// @Summary      Is Healthy Service
// @Success      200
// @Router       /healthy [get]
func IsHealthy(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "is healthy",
	})
}

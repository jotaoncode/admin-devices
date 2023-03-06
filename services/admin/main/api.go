package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "is healthy",
		})
	})
	r.GET("/device", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "here is the device detail",
		})
	})
	r.GET("/devices", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "here is the list of devices",
		})
	})
	r.POST("/device", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "adding a device",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "adding a user",
		})
	})
	r.PATCH("/user/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "updating user",
		})
	})
	r.DELETE("/user/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "removing user",
		})
	})
	r.PUT("/assign-device/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "assign a device to a user",
		})
	})
	r.PUT("/deassign-device/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "deassign a device to a user",
		})
	})
	r.GET("/user-devices/:userId", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "here the list of devices to a given user",
		})
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

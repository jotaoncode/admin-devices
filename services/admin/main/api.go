package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	gorm.Model
	Name         string
	EmployeeAbbr string
}
type Devices struct {
	gorm.Model
	MAC         string
	Name        string
	IP          string
	Description string
}
type UserDevice struct {
	gorm.Model
	UserId   uint
	DeviceId uint
}

var db *gorm.DB

func getDb() *gorm.DB {
	if db == nil {
		var host = os.Getenv("POSTGRES_SERVICE_HOST")
		dsn := "host=" + host + " user=postgres password=12345 dbname=admin port=5432 sslmode=disable"
		dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			println(err.Error())
			panic("failed to connect database")
		} else {
			println("Database connected")
			println(dbInstance.Name())
			db = dbInstance
		}
	}
	return db
}
func main() {

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "is healthy",
		})
	})
	r.GET("/device", func(c *gin.Context) {
		var device Devices
		getDb().First(&device, "1")
		c.JSON(http.StatusOK, gin.H{
			"id":          device.ID,
			"MAC":         device.MAC,
			"description": device.Description,
			"name":        device.Name,
			"created_at":  device.CreatedAt,
			"updated_at":  device.UpdatedAt,
		})
	})
	r.GET("/devices", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "here is the list of devices",
		})
	})
	r.POST("/device", func(c *gin.Context) {
		var device = &Devices{Name: "lenovo", MAC: "09-A5-4A-85-37-21", Description: "i7 13500"}
		getDb().Create(device)
		c.JSON(http.StatusOK, gin.H{
			"id": device.ID,
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

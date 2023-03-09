package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jotaoncode/greenbone/main/controllers"
)

// @title           Admin devices Example API
// @version         1.0
// @description     Administrate devices to users

// @host      localhost:9090
// @BasePath  /api/v1
func main() {
	r := gin.Default()
	// TODO add to deployment ready check
	r.GET("/healthy", controllers.IsHealthy)
	r.GET("/device/:deviceId", controllers.GetDevice)
	r.GET("/devices", controllers.GetDevices)
	r.POST("/device", controllers.CreateDevice)
	r.POST("/user", controllers.CreateUser)
	r.DELETE("/user/:userId", controllers.DeleteUser)
	r.POST("/user/:userId/device/:deviceId", controllers.AssignDeviceToUser)
	r.DELETE("/user/:userId/device/:deviceId", controllers.DeassignDeviceToUser)
	r.GET("/user/:userId/devices", controllers.GetUserDevices)
	// TODO Remove once this is defined only for testing purposes
	r.POST("/seed", controllers.CreateSeed)
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

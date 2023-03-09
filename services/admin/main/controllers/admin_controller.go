package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jotaoncode/greenbone/main/models"
	"net/http"
)

// GetDevice godoc
// @Summary      Gets a device for a given device id
// @Description  get device by id
// @Tags         devices
// @Accept       json
// @Produce      json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /device/:deviceId [get]
func GetDevice(c *gin.Context) {
	var device, err = models.GetDevice(c.Param("deviceId"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"devices": device,
	})
}

// GetDevices godoc
// @Summary      Gets all the devices
// @Description  list of devices to administrate
// @Tags         devices
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /devices [get]
func GetDevices(c *gin.Context) {
	var devices, err = models.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"devices": devices,
	})
}

// CreateDevice godoc
// @Summary      Create a device
// @Description  creates a new device
// @Tags         devices
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /device [post]
func CreateDevice(c *gin.Context) {
	var device models.Devices
	c.BindJSON(&device)
	var deviceId, err = models.CreateDevice(device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deviceId": deviceId})
}

// CreateUser godoc
// @Summary      Create a user
// @Description  creates a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /user [post]
func CreateUser(c *gin.Context) {
	var user models.Users
	c.BindJSON(&user)
	var userId, err = models.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deviceId": userId})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /user/:userId [delete]
func DeleteUser(c *gin.Context) {
	var userId = c.Param("userId")
	var _, err = models.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// AssignDeviceToUser godoc
// @Summary      Assigns a device to a user
// @Description  Assigns a device to a user, informing in case 3 or more assigned devices to user
// @Tags         userdevice
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /user/:userId/device/:deviceId [post]
func AssignDeviceToUser(c *gin.Context) {
	var userId = c.Param("userId")
	var deviceId = c.Param("deviceId")
	var _, err = models.AssignUserDevice(userId, deviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// DeassignDeviceToUser godoc
// @Summary      Deassign device to a user
// @Description  Deassign device to a user by id
// @Tags         userdevice
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /user/:userId/device/:deviceId [delete]
func DeassignDeviceToUser(c *gin.Context) {
	var userId = c.Param("userId")
	var deviceId = c.Param("deviceId")
	var _, err = models.DeassignDeviceToUser(userId, deviceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// GetUserDevices godoc
// @Summary      Ge user devices
// @Description  a list of devices assigned to the user
// @Tags         userdevice
// @Accept       json
// @Produce      json
// @Success      200
//
//	@Failure      500
//
// @Router       /user/:userId/devices [get]
func GetUserDevices(c *gin.Context) {
	var userId = c.Param("userId")
	var userDevices, err = models.GetUserDevices(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userDevices)
}

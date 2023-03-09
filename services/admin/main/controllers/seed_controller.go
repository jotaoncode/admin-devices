package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jotaoncode/greenbone/main/connectors"
	"github.com/jotaoncode/greenbone/main/models"
	"net/http"
)

// TODO remove or provide just for testing purposes
func CreateSeed(c *gin.Context) {
	var assignOnce = uuid.New().String()
	var assignTwoTimes1 = uuid.New().String()
	var assignTwoTimes2 = uuid.New().String()
	var devices = &[]models.Devices{
		{ID: assignOnce, Name: "lenovo", MAC: "09-A5-4A-85-37-21", Description: "i7 13500"},
		{ID: assignTwoTimes1, Name: "pro", MAC: "09-A6-5A-85-37-21", Description: "i5 13500"},
		{ID: assignTwoTimes2, Name: "unnasigned lenovo", MAC: "09-B5-4A-85-37-21", Description: "i7 12500"},
		{ID: uuid.New().String(), Name: "unnasigned pro", MAC: "09-B6-5A-85-37-21", Description: "i5 12500"},
		{ID: uuid.New().String(), Name: "unnasigned air", MAC: "04-B3-5B-85-37-21", Description: "i3 11500"},
	}
	connectors.GetDb().Create(devices)
	var assignedUserOnce = uuid.New().String()
	var assignedUserTwoTimes = uuid.New().String()
	var users = &[]models.Users{
		{ID: assignedUserOnce, Name: "Juan Jose Garcia", Abbr: "JJG"},
		{ID: assignedUserTwoTimes, Name: "Paula Perez", Abbr: "PPE"},
		{ID: uuid.New().String(), Name: "Max Musterman", Abbr: "MMU"},
		{ID: uuid.New().String(), Name: "John Doe", Abbr: "JDO"},
		{ID: uuid.New().String(), Name: "Alejandra Ramirez", Abbr: "ARA"},
	}
	connectors.GetDb().Create(users)
	var assignations = &[]models.UserDevices{
		{ID: uuid.New().String(), UserId: assignedUserOnce, DeviceId: assignOnce},
		{ID: uuid.New().String(), UserId: assignedUserTwoTimes, DeviceId: assignTwoTimes1},
		{ID: uuid.New().String(), UserId: assignedUserTwoTimes, DeviceId: assignTwoTimes2},
	}
	connectors.GetDb().Create(assignations)

	c.JSON(http.StatusOK, gin.H{
		"assignedUserOnce":     assignedUserOnce,
		"assignedUserTwoTimes": assignedUserTwoTimes,
		"deviceAssignedOnce":   assignTwoTimes1,
		"deviceAssignedTwice":  assignTwoTimes2,
	})
}

package models

import (
	"github.com/jotaoncode/greenbone/main/connectors"
)

func GetUserDevices(userId string) ([]UserDevice, error) {
	var userDevices []UserDevices
	connectors.GetDb().Find(&userDevices, "user_id = ?", userId)
	var deviceIds []string

	for i := 0; i < len(userDevices); i++ {
		deviceIds = append(deviceIds, userDevices[i].DeviceId)
	}
	var devices []Devices
	connectors.GetDb().Find(&devices, "id IN ?", deviceIds)
	var user Users
	connectors.GetDb().First(&user, "id = ?", userId)
	//fmt.Printf("%+v\n", usersAssigned)
	var assignedUserDevices []UserDevice
	for i := range deviceIds {
		var userDevice UserDevice
		userDevice.UserId = userId
		userDevice.DeviceId = deviceIds[i]
		userDevice.Name = user.Name
		userDevice.Abbr = user.Abbr

		for device := range devices {
			if devices[device].ID == deviceIds[i] {
				userDevice.DeviceDescription = devices[device].Description
				break
			}
		}
		assignedUserDevices = append(assignedUserDevices, userDevice)
	}
	return assignedUserDevices, nil
}

func CreateDevice(device Devices) (string, error) {
	deviceCreated := connectors.GetDb().Create(&device)

	return device.ID, deviceCreated.Error
}

func GetDevice(deviceId string) (Devices, error) {
	var device Devices
	queryResult := connectors.GetDb().First(&device, deviceId)
	return device, queryResult.Error
}

func GetDevices() (Devices, error) {
	var devices Devices
	queryResult := connectors.GetDb().Find(&devices)
	return devices, queryResult.Error
}

func CreateUser(user Users) (string, error) {
	queryResult := connectors.GetDb().Create(user)
	return user.ID, queryResult.Error
}

func DeleteUser(userId string) (string, error) {
	queryResult := connectors.GetDb().Delete(&Users{}, userId)
	return userId, queryResult.Error
}

func AssignUserDevice(userId string, deviceId string) (string, error) {
	var userDevice = UserDevices{UserId: userId, DeviceId: deviceId}
	queryResult := connectors.GetDb().Create(&userDevice)
	var userAssignedDevices []UserDevices
	connectors.GetDb().Find(userAssignedDevices, "user_id", userId)
	if len(userAssignedDevices) > 2 {
		println("Inform about it")
	}
	return userDevice.ID, queryResult.Error
}

func DeassignDeviceToUser(userId string, deviceId string) (string, error) {
	var userDevice = UserDevices{UserId: userId, DeviceId: deviceId}
	queryResult := connectors.GetDb().Delete(&userDevice)
	return userDevice.ID, queryResult.Error
}

package models

import (
	"github.com/google/uuid"
	"github.com/jotaoncode/greenbone/main/connectors"
)

type AdminModel struct{}

func (adminModel AdminModel) GetUserDevices(userId string) ([]UserDevice, error) {
	var userDevices []UserDevices
	connectors.DB{}.GetDb().Find(&userDevices, "user_id = ?", userId)
	var deviceIds []string

	for i := 0; i < len(userDevices); i++ {
		deviceIds = append(deviceIds, userDevices[i].DeviceId)
	}
	var devices []Devices
	connectors.DB{}.GetDb().Find(&devices, "id IN ?", deviceIds)
	var user Users
	connectors.DB{}.GetDb().First(&user, "id = ?", userId)
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

func (adminModel AdminModel) CreateDevice(deviceCreationRequest CreateDeviceRequest) (string, error) {
	var device = &Devices{
		ID:          uuid.New().String(),
		Name:        deviceCreationRequest.Name,
		MAC:         deviceCreationRequest.MAC,
		Description: deviceCreationRequest.Description,
		IP:          deviceCreationRequest.IP,
	}
	deviceCreated := connectors.DB{}.GetDb().Create(&device)

	return device.ID, deviceCreated.Error
}

func (adminModel AdminModel) GetDevice(deviceId string) (Devices, error) {
	var device Devices
	queryResult := connectors.DB{}.GetDb().First(&device, deviceId)

	return device, queryResult.Error
}

func (adminModel AdminModel) GetDevices() (Devices, error) {
	var devices Devices
	queryResult := connectors.DB{}.GetDb().Find(&devices)

	return devices, queryResult.Error
}

func (adminModel AdminModel) CreateUser(userCreationRequest CreateUserRequest) (string, error) {
	// TODO from name to abbr
	var user = &Users{ID: uuid.New().String(), Name: userCreationRequest.Name, Abbr: "MMU"}
	queryResult := connectors.DB{}.GetDb().Create(&user)

	return user.ID, queryResult.Error
}

func (adminModel AdminModel) DeleteUser(userId string) (string, error) {
	queryResult := connectors.DB{}.GetDb().Delete(&Users{}, userId)

	return userId, queryResult.Error
}

func (adminModel AdminModel) AssignUserDevice(userId string, deviceId string) (string, error) {
	var userDevice = UserDevices{ID: uuid.New().String(), UserId: userId, DeviceId: deviceId}
	queryResult := connectors.DB{}.GetDb().Create(&userDevice)
	var userAssignedDevices []UserDevices
	connectors.DB{}.GetDb().Find(&userAssignedDevices, "user_id", userId)
	if len(userAssignedDevices) > 2 {
		var user Users
		connectors.DB{}.GetDb().First(&user, "id = ?", userId)
		_, err := connectors.NotifyExcessOfDevices(user.Abbr, "warning")
		if err != nil {
			return userDevice.ID, err
		}
	}
	return userDevice.ID, queryResult.Error
}

func (adminModel AdminModel) DeassignDeviceToUser(userId string, deviceId string) (string, error) {
	var userDevice = UserDevices{UserId: userId, DeviceId: deviceId}
	queryResult := connectors.DB{}.GetDb().Delete(&userDevice)
	return userDevice.ID, queryResult.Error
}

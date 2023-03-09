package models

import (
	"github.com/jotaoncode/greenbone/main/connectors"
	"testing"
)

// TODO Improve using stubs spies and more granular testing
// Test Assign User Device, it should create a new registry in db
func TestAssignUserDevice(t *testing.T) {
	adminModel := &AdminModel{}

	userCreationRequest := CreateUserRequest{Name: "Juan Garcia"}

	userId, errCreateUser := adminModel.CreateUser(userCreationRequest)
	if errCreateUser != nil {
		t.Fatalf("Failed to create user")
	}

	deviceCreationRequest := CreateDeviceRequest{Name: "i7 12500", Description: "Computer for analytics", MAC: "MACADDRESS", IP: "192.168.172.1"}
	deviceId, errCreateDevice := adminModel.CreateDevice(deviceCreationRequest)

	if errCreateDevice != nil {
		t.Fatalf("Failed to create device")
	}
	userDeviceId, errAssignDevice := adminModel.AssignUserDevice(userId, deviceId)

	if errAssignDevice != nil {
		t.Fatalf("Failed to assign device to user")
	}
	userDevice := UserDevice{}
	connectors.DB{}.GetDb().First(&userDevice, "id = ?", userDeviceId)

	if userDevice.UserId != userId {
		t.Fatalf("Device does not belong to user")
	}

	if userDevice.DeviceId != deviceId {
		t.Fatalf("Device wrongly assigned")
	}
}

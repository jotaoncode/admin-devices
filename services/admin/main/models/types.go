package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID   string
	Name string
	Abbr string
}
type Devices struct {
	gorm.Model
	ID          string
	MAC         string
	Name        string
	IP          string
	Description string
}
type UserDevices struct {
	gorm.Model
	ID       string
	UserId   string
	DeviceId string
}

type UserDevice struct {
	UserId            string `json:"userId"`
	DeviceId          string `json:"deviceId"`
	DeviceDescription string `json:"deviceDescription"`
	Name              string `json:"name"`
	Abbr              string `json:"abbr"`
}

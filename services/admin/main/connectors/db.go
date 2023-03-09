package connectors

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func GetDb() *gorm.DB {
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

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	return connect()
}

func connect () *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
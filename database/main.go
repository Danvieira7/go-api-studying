package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Create and use interface for abstract DB.
func GetDB() *gorm.DB {
	return connect()
}

// Make the address of SQLite DB come from an ENV variable.
func connect () *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test_database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
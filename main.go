package main

import (
	"github.com/hoisie/web"
	"github.com/soufraz/go-api-study/database"
	"github.com/soufraz/go-api-study/users"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		return nil
	}

	return db
}

func main () {
	migrate(database.GetDB())

	web.Get("/users", users.Index)
	web.Post("/users", users.Create)
	web.Get("/users/(.*)", users.Get)
	web.Run("0.0.0.0:9977")
}

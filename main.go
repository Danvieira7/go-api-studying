package main

import (
	"github.com/hoisie/web"
	"gorm.io/gorm"

	"github.com/Danvieira7/go-api-studying/database"
	"github.com/Danvieira7/go-api-studying/users"
)

func migrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		return nil
	}

	return db
}

func main() {
	migrate(database.GetDB())

	web.Get("/users", users.Index)
	web.Get("/users/(.*)", users.Get)
	web.Delete("/users/(.*)", users.Delete)
	web.Put("/users/(.*)", users.Edit)
	web.Post("/users", users.Create)
	
	web.Run("0.0.0.0:9977")
}

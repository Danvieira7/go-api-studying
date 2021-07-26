package main

import (
	"github.com/hoisie/web"
	"gorm.io/gorm"

	"github.com/Danvieira7/go-api-studying/database"
	"github.com/Danvieira7/go-api-studying/users"
)

// Find out how to move this migrate method to the database package.  
// If we move this migrate method to database package and try 
// calling it here we get a circular dependency error.
func migrate(db *gorm.DB) *gorm.DB {
	err := db.AutoMigrate(&users.User{})
	if err != nil {
		return nil
	}

	return db
}


func main() {
	migrate(database.GetDB())

	// Figure ou how to make dependency inversion from db into routes methods.
	// The idea is to decouple db fom routes methods in order to get
	// better unit testing capability.
	web.Get("/users", users.Index)
	web.Get("/users/(.*)", users.Get)
	web.Delete("/users/(.*)", users.Delete)
	web.Put("/users/(.*)", users.Edit)
	web.Post("/users", users.Create)
	// Does it make sense to move this routes to an independent route file?
	
	// This address should come from an ENV variable.
	web.Run("0.0.0.0:9977")
}

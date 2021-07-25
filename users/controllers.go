package users

import (
	"encoding/json"

	"github.com/Danvieira7/go-api-studying/database"
	"github.com/hoisie/web"
)

func Index(ctx *web.Context) string {
	// Figure out how to set this line once to be reusable.
	ctx.SetHeader("Content-Type", "application/json", true)

	var users []User
	
	// Turn this db variable into a dependency injection.
	db := database.GetDB()
	db.Find(&users)

	// Figure out if we can turn this into reusable pattern return for Json OBJ.
	// Handle error. 
	bytes, _ := json.Marshal(users)

	return string(bytes)
}

func Get(ctx *web.Context, id string) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	var user User

	db := database.GetDB()
	db.Find(&user, id)

	bytes, _ := json.Marshal(user)

	return string(bytes)
}

// Validate required fields and type of fields (all persistency methods).
func Create(ctx *web.Context) string {	
	ctx.SetHeader("Content-Type", "application/json", true)

	user := &User{}
	json.NewDecoder(ctx.Request.Body).Decode(&user)

	db := database.GetDB()
	db.Create(user)

	bytes, _ := json.Marshal(user)
	return string(bytes)
}

func Delete(ctx *web.Context, id string) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	db := database.GetDB()
	db.Delete(&User{}, id)

	result := map[string]string{"msg": "User successfully deleted."}

	bytes, _ := json.Marshal(result)

	return string(bytes)
}

func Edit(ctx *web.Context, id string) string {
	ctx.SetHeader("Content-Type", "application/json", true)
		
	var user User

	db := database.GetDB()
	db.Find(&user, id)
	
	json.NewDecoder(ctx.Request.Body).Decode(&user)
	db.Save(&user)

	bytes, _ := json.Marshal(user)

	return string(bytes)
}

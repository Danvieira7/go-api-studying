package users

import (
	"encoding/json"
	"github.com/hoisie/web"
	"github.com/soufraz/go-api-study/database"
)

var userJackMocked = map[string]string{
	"name": "Jack",
	"age": "24",
}

func Index (ctx *web.Context) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	collectionResult := []map[string]string{
		userJackMocked,
		userJackMocked,
		userJackMocked,
		userJackMocked,
	}

	bytes, _ := json.Marshal(collectionResult)

	return string(bytes)
}

func Get (ctx *web.Context, id string) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	bytes, _ := json.Marshal(userJackMocked)

	return string(bytes)
}

func Create (ctx *web.Context) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	user := &User{}
	json.NewDecoder(ctx.Request.Body).Decode(&user)

	db := database.GetDB()
	db.Create(user)

	bytes, _ := json.Marshal(user)
	return string(bytes)
}
package users

import (
	"encoding/json"
	"github.com/hoisie/web"
)

var userJackMocked = map[string]string{
	"name": "Jack",
	"age": "24",
}

func Index (ctx *web.Context, val string) string {
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

func Get (ctx *web.Context, val string) string {
	ctx.SetHeader("Content-Type", "application/json", true)

	bytes, _ := json.Marshal(userJackMocked)

	return string(bytes)
}
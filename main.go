package main

import (
	"github.com/hoisie/web"
	"github.com/soufraz/go-api-study/users"
)

func main () {
	web.Get("/users(.*)", users.Index)
	web.Get("/user(.*)", users.Get)
	web.Run("0.0.0.0:9977")
}

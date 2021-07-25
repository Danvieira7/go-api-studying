package main

import "github.com/hoisie/web"

func index(val string) string {
	return "hello " + val
}

func main () {
	web.Get("/(.*)", index)
	web.Run("0.0.0.0:9977")
}

package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee!</h1>")
	})
	engine.GET("/hello/:name", func(ctx *gee.Context) {
		ctx.String(http.StatusOK, "Hello Gee!string!!")
	})
	engine.Run(":8888")
}

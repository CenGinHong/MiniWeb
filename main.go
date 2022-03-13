package main

import (
	"MiniWeb/web"
	"net/http"
)

func main() {
	r := web.New()
	r.GET("/", func(ctx *web.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(ctx *web.Context) {
		// expect /hello?name=geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})
	r.GET("/hello/:name", func(ctx *web.Context) {
		// expect /hello/geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Param("name"), ctx.Path)
	})
	r.GET("/assets/*filepath", func(ctx *web.Context) {
		ctx.JSON(http.StatusOK, web.H{"filepath": ctx.Param("filepath")})
	})
	_ = r.Run(":9999")
}

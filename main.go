package main

import (
	"MiniWeb/web"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	r := web.New()
	r.GET("/", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "URL.Path = %q\n", ctx.Path)
	})
	r.GET("/hello", func(ctx *web.Context) {
		sb := strings.Builder{}
		for k, v := range ctx.Req.Header {
			sb.WriteString(fmt.Sprintf("Header[%q] = %q\n", k, v))
		}
		ctx.String(http.StatusOK, sb.String())
	})
	_ = r.Run(":9999")
}

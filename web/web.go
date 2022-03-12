package web

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRouter 添加路由表，例如key是GET-/  POST-/hello
func (e *Engine) addRouter(method string, pattern string, handler HandleFunc) {
	e.router.addRouter(method, pattern, handler)
}

// GET 将该方法注册到路由表中
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST 将该方法注册到路由表中
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRouter("POST", pattern, handler)
}

// Run 交由Engine处理http
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.router.handle(c)
}

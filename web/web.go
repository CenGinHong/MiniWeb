package web

import (
	"fmt"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// addRouter 添加路由表，例如key是GET-/  POST-/hello
func (e *Engine) addRouter(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
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
	// 组成key
	key := req.Method + "-" + req.URL.Path
	// 查表取出构成函数
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		// 没有找到处理方法
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

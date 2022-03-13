package web

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandleFunc),
	}
}

// parsePattern 解析路由成数组，且只允许出现一个*
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRoute
func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	// 得到拆分路径
	parts := parsePattern(pattern)
	// 组成key值
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRouter
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	// 根据method分类
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	// 在字典树下搜索node
	n := root.search(searchParts, 0)
	// 如果n存在
	if n != nil {
		// 拆分n.pattern，注意这里可能拆出来是 /p/:lang/go
		parts := parsePattern(n.pattern)
		// parts
		for index, part := range parts {
			// 如果这个是匹配值
			if part[0] == ':' {
				// 取出来作为参数，即将：id映射到一个具体值
				params[part[1:]] = searchParts[index]
			}
			// 如果是通配符
			if part[0] == '*' && len(part) > 1 {
				// 将后面的全部通配进去
				params[part[1:]] = strings.Join(searchParts[index:], "/")
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	// 查找存不存在该节点，并且进行参数映射
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		// 存在，把路径参数映射到上下文中
		c.Params = params
		key := c.Method + "-" + n.pattern
		// 执行handler
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

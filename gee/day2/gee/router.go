package gee

import "net/http"

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if hander, ok := r.handlers[key]; ok {
		hander(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %q", c.Path)
	}
}

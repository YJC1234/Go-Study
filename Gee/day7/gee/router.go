package gee

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandleFunc
}

//roots key eg, roots['GET'] roots['POST']
//handlers key eg, handlers['GET-/p/:lang/doc'] handlers['POST-/p/book']

func newRouter() *router {
	return &router{roots: make(map[string]*node), handlers: make(map[string]HandleFunc)}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler

	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
}

// 除了返回找到的node，还返回map指明参数(:lang)
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	_, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	node := r.roots[method].search(searchParts, 0)
	if node != nil {
		parts := parsePattern(node.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node, params
	}
	return nil, nil

}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)

	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.hanlders = append(c.hanlders, r.handlers[key])
	} else {
		c.hanlders = append(c.hanlders, func(ctx *Context) {
			ctx.String(http.StatusNotFound, "404 NOT FOUND: %q", c.Path)
		})
	}
	c.Next()
}

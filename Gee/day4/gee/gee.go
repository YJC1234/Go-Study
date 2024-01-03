package gee

import (
	"net/http"
)

// HandleFunc is the request hander
type HandleFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	groups []*RouterGroup
}

// New is constructor of gee.Engine
func New() *Engine {
	group := newRootGroup()
	return &Engine{
		RouterGroup: group,
		groups:      []*RouterGroup{group},
	}
}

// Group engine的层面上创建新分组
func (e *Engine) Group(prefix string) *RouterGroup {
	group := e.RouterGroup.Group(prefix)
	e.groups = append(e.groups, group)
	return group
}

// Run : start a http server
func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	e.handle(c)
}

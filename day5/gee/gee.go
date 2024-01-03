package gee

import (
	"net/http"
	"strings"
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
	engine := &Engine{}
	group := newRootGroup(engine)
	engine.RouterGroup = group
	engine.groups = []*RouterGroup{group}
	return engine
}

// Run : start a http server
func (e *Engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(c.Path, group.prefix){
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c.hanlders = middlewares
	e.handle(c)
}

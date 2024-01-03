package gee

import (
	"net/http"
)

// HandleFunc is the request hander
type HandleFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is constructor of gee.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET : add a get request
func (engine *Engine) GET(pattern string, hander HandleFunc) {
	engine.addRoute("GET", pattern, hander)
}

// POST : add a post request
func (engine *Engine) POST(pattern string, hander HandleFunc) {
	engine.addRoute("POST", pattern, hander)
}

// Run : start a http server
func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

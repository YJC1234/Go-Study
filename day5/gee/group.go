package gee

import (
	"log"
)

// RouterGroup 职责是对路由进行分组(对注册的路由添加组前缀)，在组的粒度上对路由进行控制
type RouterGroup struct {
	engine *Engine //共享实例，提供访问接口
	*router
	prefix      string
	parent      *RouterGroup
	middlewares []HandleFunc
}

// 每一个分组必定有根，因此可以创建一个根路由分组,新的分组必然通过根路由分组添加
func newRootGroup(engine *Engine) *RouterGroup {
	return &RouterGroup{
		prefix: "",
		router: newRouter(),
		engine: engine}
}

// Group 添加子分组,newPrefix为子分组的前缀
func (group *RouterGroup) Group(newPrefix string) *RouterGroup {
	newPrefix = group.prefix + newPrefix
	newRouterGroup := &RouterGroup{
		router: group.router,
		prefix: newPrefix,
		parent: group,
		engine: group.engine,
		//middlewares: group.middlewares,
	}
	group.engine.groups = append(group.engine.groups, newRouterGroup)
	return newRouterGroup
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandleFunc) {
	pattern = group.prefix + pattern
	log.Printf("Route %4s - %s", method, pattern)
	group.router.addRoute(method, pattern, handler)
}

// GET add get request router
func (group *RouterGroup) GET(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST add post request router
func (group *RouterGroup) POST(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

// Use 使用中间件
func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

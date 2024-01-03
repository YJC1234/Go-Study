package gee

import (
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/hello/:name"), []string{"hello", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/hello/*filename/dinner"), []string{"hello", "*filename"})

	if !ok {
		t.Fatal("test ParsePattern fail")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, params := r.getRoute("GET", "/hello/b")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/hello/:name" {
		t.Fatal("pattern error!")
	}
	if params["name"] != "b" {
		t.Fatal("params error!")
	}
}

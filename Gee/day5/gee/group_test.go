package gee

import (
	"testing"
)

func TestGroup(t *testing.T) {
	e := New()
	v1 := e.Group("/v1")
	v2 := v1.Group("/v2")
	v3 := v2.Group("/v3")
	if v3.prefix != "/v1/v2/v3" {
		t.Fatal("group error!")
	}
}

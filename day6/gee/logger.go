package gee

import (
	"log"
	"time"
)

// Logger 日志，记录时间
func Logger() HandleFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

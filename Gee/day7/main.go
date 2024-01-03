package main

import (
	"gee"
	"net/http"
)

func main() {
	e := gee.Default()

	e.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	e.GET("/panic", func(c *gee.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	e.Run(":8888")
}

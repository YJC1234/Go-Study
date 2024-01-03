package main

import (
	"fmt"
	"gee"
	"net/http"
	"text/template"
	"time"
)

type student struct {
	Name string
	Age  int8
}

// FormatAsDate 格式化日期
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	e := gee.New()
	e.Use(gee.Logger())
	e.SetFuncMap(template.FuncMap{
		"FormatAsData": FormatAsDate,
	})
	e.LoadHTMLGlob("templates/*")
	//e.Static("/assets", "./static")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}

	e.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	e.Run(":8888")
}

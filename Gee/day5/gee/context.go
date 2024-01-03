package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H Header结构简写
type H map[string]interface{}

// Context 请求相关信息集合
type Context struct {
	//origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Method     string
	Path       string
	Params     map[string]string
	StatusCode int
	//middleware
	hanlders []HandleFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

// Param 返回c.Params中key对应value
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 返回表单key对应string
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 返回query对应string
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置请求码后发送请求头
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 构造json相应
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 构造data响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 构造HTML响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}

//Next 中间件调用(便于前后处理)
func (c *Context) Next() {
	c.index++
	for ;c.index<len(c.hanlders);c.index++{
		c.hanlders[c.index](c)
	}
}
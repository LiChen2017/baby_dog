package main

import (
	"net/http"
	"html/template"
	"github.com/gin-gonic/gin"
)

func handler_index1(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"NAME": "Chen Li",
		"address": "<a href='http://www.w3school.com.cn'> W3School </a>",
	})
}
func handler_index2(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"NAME": "Li Chen",
		"address": "<a href='https://liwenzhou.com'> BLOG </a>",
	})
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"safe": func(val string) template.HTML{
			return template.HTML(val)
		},
	})
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts", handler_index1)
	r.GET("/users", handler_index2)

	r.Run(":9090")
}
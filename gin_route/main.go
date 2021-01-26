package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func handler1(c *gin.Context) {
	// upload a file
	c.HTML(http.StatusOK, "index.html", nil)
}
func handler2(c *gin.Context) {
	// download a file
	file, err := c.FormFile("f")
	if err == nil {
		c.SaveUploadedFile(file, "new-file")
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
func handler3(c *gin.Context) {
	// re-direct destination
	c.Redirect(http.StatusMovedPermanently, "https://liwenzhou.com")
}
func handler4(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "this is handler 4",
	})
}
func handler5(c *gin.Context) {
	// curl -X GET "localhost:9090/route"
	c.JSON(http.StatusOK, gin.H{
		"type": "GET",
	})
}
func handler6(c *gin.Context) {
	// curl -X POST "localhost:9090/route"
	c.JSON(http.StatusOK, gin.H{
		"type": "POST",
	})
}
func handler7(c *gin.Context) {
	// curl -X GET/POST/DELETE/CONNECT/TRACE "localhost:9090/route"
	switch c.Request.Method {
	case http.MethodGet:
		c.JSON(http.StatusOK, gin.H{"type":"any-GET",})
	case http.MethodPost:
		c.JSON(http.StatusOK, gin.H{"type":"any-POST",})
	case http.MethodDelete:
		c.JSON(http.StatusOK, gin.H{"type":"any-DELETE",})
	default:
		c.JSON(http.StatusBadRequest, gin.H{"message":"invalid type",})
	}
}
func handler8(c *gin.Context) {
	// http://127.0.0.1:9090/video/home
	c.JSON(http.StatusOK, gin.H{"group":"video",})
}
func handler9(c *gin.Context) {
	// http://127.0.0.1:9090/shop/home
	c.JSON(http.StatusOK, gin.H{"group":"shop",})
}
func handler10(c *gin.Context) {
	// http://127.0.0.1:9090/dfljslls
	c.JSON(http.StatusNotFound, gin.H{"message": "empty page",})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")
	r.GET("/upload", handler1)
	r.POST("/download", handler2)
	r.GET("/redirect", handler3)
	r.GET("/route1", handler4)
	r.GET("/route2", func(c *gin.Context) {
		// re-direct route
		c.Request.URL.Path = "/route1"
		r.HandleContext(c)
	})
	r.GET("/route", handler5)
	r.POST("/route", handler6)
	r.Any("/route_any", handler7)
	groupV := r.Group("/video")
	{
		groupV.GET("/home", handler8)
	}
	groupS := r.Group("/shop")
	{
		groupS.GET("/home", handler9)
	}

	r.NoRoute(handler10)
	r.Run(":9090")
}
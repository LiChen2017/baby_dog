package main

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

func handler0(c *gin.Context) {
	stamp := time.Now()
	fmt.Println("...... get in handler0 ......")
	//c.Abort()
	//return
	c.Set("year", 2021)
	c.Next()
	fmt.Println("...... get out handler0 ......", time.Since(stamp))
}
func handler1(c *gin.Context) {
	fmt.Println("...... get in handler1 ......")
	val, _ := c.Get("year")
	c.JSON(http.StatusOK, gin.H{"message": "ok", "year": val})
	fmt.Println("...... get out handler1 ......")
}

func main() {
	r := gin.Default()
	r.GET("/midware", handler0, handler1)

	r.Run(":9090")
}
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Describe struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
type Account struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func handler1(c *gin.Context) {
	// return map, gin.H, struct in JSON
	var person = map[string]interface{}{"name": "Chen Li",}
	var info = gin.H{"age": 3, "gender": "female",}
	var introduction = Describe{"pang pang", 3}
	c.JSON(http.StatusOK, gin.H{
		"name": person,
		"detail": info,
		"dog": introduction,
	})
}
func handler2(c *gin.Context) {
	// extract request parameter: http://127.0.0.1:9090/query?info1=aaa&info2=bbb&info3=ccc
	var info1, info2, info3 string
	var ok bool = true
	info1 = c.Query("info1")
	info2 = c.DefaultQuery("info2", "use_default_message")
	info3, ok = c.GetQuery("info3")
	if ok {
		c.String(http.StatusOK, "get result: " + info1 + ", " + info2 + ", " + info3)
	} else {
		c.String(http.StatusOK, "empty result")
	}
}
func handler3(c *gin.Context) {
	// http://127.0.0.1:9090/login
	c.HTML(http.StatusOK, "login.html", nil)
}
func handler4(c *gin.Context) {
	var info1 = c.DefaultPostForm("username", "use_default_name")
	var info2 = c.DefaultPostForm("password", "use_default_password")
	var info3 = c.DefaultPostForm("passworddd", "use_default_passworddd")
	c.HTML(http.StatusOK, "temp.html", gin.H{
		"name": info1,
		"password": info2,
		"passworddd": info3,
	})
}
func handler5(c *gin.Context) {
	// http://127.0.0.1:9090/input/china/hangzhou
	c.JSON(http.StatusOK, gin.H{
		"value1": c.Param("param1"),
		"value2": c.Param("param2"),
	})
}
func handler6(c *gin.Context) {
	// curl -X POST "localhost:9090/struct?username=kkk&password=iiilll"
	// curl -H "Content-Type:application/json" -X POST -d '{"username":"qqqq","password":"pppp"}' "localhost:9090/struct"
	var info Account
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": info,
		})
	}
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("login.html", "temp.html")
	r.GET("/index", handler1)
	r.GET("/query", handler2)
	r.GET("/login", handler3)
	r.POST("/login", handler4)
	r.GET("/input/:param1/:param2", handler5)
	r.POST("/struct", handler6)

	r.Run(":9090")
}
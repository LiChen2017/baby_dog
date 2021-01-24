package main

import (
    "os"
    "io"
    //"log"
    //"time"
    "bytes"
    //"context"
    //"syscall"
    "net/http"
    "io/ioutil"
    //"os/signal"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/autotls"
)

type Person struct {
    Name string `form:"name" binding:"required"`
    Age int `form:"age" binding:"required,gt=1"`
}

func main() {
    r := gin.Default()

    // TEST 1
    /*handler := func(c *gin.Context) {c.JSON(200, gin.H{"message":"ping_get",})}
    r.GET("/ping_get", handler)
    handler = func(c *gin.Context) {c.String(200, "ping_post")}
    r.POST("/ping_post", handler)
    handler = func(c *gin.Context) {c.String(200, "ping_delete")}
    r.Handle("DELETE", "/ping_delete", handler)
    handler = func(c *gin.Context) {c.String(200, "ping_test")}
    r.Handle("GET", "/ping_test", handler)
    handler = func(c *gin.Context) {c.JSON(200, gin.H{"name":"ANY", "age":3,})}
    r.Any("/ping_any", handler)*/
    // test ping_get/ping_post/ping_delete
    // curl -X GET "localhost:8080/ping_get"
    // curl -X POST "localhost:8080/ping_post"
    // curl -X DELETE "localhost:8080/ping_delete"
    // curl -X GET "localhost:8080/ping_test"
    // curl -X GET/POST/DELETE/CONNECT/... "localhost:8080/ping_any"

    // TEST 2
    r.Static("/test_route1", "route1")
    r.StaticFS("/test_route2", http.Dir("route2"))
    r.StaticFile("/test_route3", "route3/pic.ico")

    // TEST 3: get parameter
    handler := func(c *gin.Context) {c.JSON(200, gin.H{"Name":c.Param("NAME"), "Age":c.Param("AGE"),})}
    r.GET("information/:NAME/:AGE", handler)
    handler = func(c *gin.Context) {c.String(200, "action_test")}
    r.GET("/user/*action", handler)
    handler = func(c *gin.Context) {
	    fName := c.Query("first_name")
	    lName := c.DefaultQuery("last_name", "use_default_lname")
	    c.String(http.StatusOK, "Fname: %s, Lname: %s", fName, lName)
    }
    r.GET("/get_parameter", handler)
    handler = func(c *gin.Context) {
	    _bytes, err := ioutil.ReadAll(c.Request.Body)
	    if err != nil {
		    c.String(http.StatusBadRequest, err.Error())
		    c.Abort()
	    }
	    c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(_bytes))
	    fName := c.PostForm("first_name")
	    lName := c.DefaultPostForm("last_name", "use_default_nale")
	    c.String(http.StatusOK, "Bytes: %s, Fname: %s, Lname: %s", string(_bytes), fName, lName)
    }
    r.POST("/get_body", handler)
    handler = func(c *gin.Context) {
	    var per Person
	    if err := c.ShouldBind(&per); err == nil {
		    c.String(200, "%v", per)
	    } else {
		    c.String(200, "bind error: %v", err)
	    }
    }
    r.GET("/get_struct", handler)
    r.POST("/get_struct", handler)

    // TEST 4: verify parameter
    handler = func(c *gin.Context) {
	    var per Person
	    if err := c.ShouldBind(&per); err != nil {
		    c.String(500, "%v", err)
	    }
	    c.String(200, "%v", per)
    }
    r.GET("/verify_parameter", handler)

    // TEST 5
    handler = func(c *gin.Context) {
	    name := c.DefaultQuery("name", "use_default_name")
	    c.String(http.StatusOK, "%v", name)
    }
    fout, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(fout)
    r = gin.New()
    r.Use(gin.Logger())
    r.GET("/middle", handler)

    // TEST 6, graceful shutdown
    /*r = gin.Default()
    handler = func(c *gin.Context) {
	    time.Sleep(15 * time.Second)
	    c.String(200, "test_shutdown")
    }
    r.GET("/shutdown", handler)
    srv := &http.Server{Addr:":8085", Handler: r}
    go func() {
	    if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		    log.Fatalf("listen: %s\n", err)
	    }
    }()
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <- quit
    log.Println("! ! ! Shuting down server ! ! !")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
	    log.Fatal("! ! ! server failed shut down . . .", err)
    }
    log.Println("! ! ! server is shut down . . .")*/

    // TEST 7
    r.LoadHTMLGlob("temp/*")
    handler = func(c *gin.Context) {
	    c.HTML(200, "tmp.html", gin.H{"title":"test_template",})
    }
    r.GET("/template", handler)

    handler = func(c *gin.Context) {
	    c.String(200, "test_certificate")
    }
    r.GET("/certificate", handler)
    autotls.Run(r, "www.itpp.tk")

    r.Run()
}

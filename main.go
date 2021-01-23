package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

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
    r.Static("/test_route1", "route1/")
    r.StaticFS("/test_route2", http.Dir("route2/"))
    r.StaticFile("/test_route3", "./route3/pic.ico")

    // TEST 3
    handler_url := func(c *gin.Context) {c.JSON(200, gin.H{"Name":c.Param("NAME"), "Age":c.Param("AGE"),})}
    r.GET("information/:NAME/:AGE", handler_url)

    r.Run()
}

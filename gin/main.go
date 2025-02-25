package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world!")
	})
	// 参数路由，路径参数
	server.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)
	})
	// 查询参数，http://localhost:8080/order?id=123
	server.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "ID是, "+id)
	})
	// 通配符路由，http://localhost:8080/views/hello.html
	server.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html")
		ctx.String(http.StatusOK, "view是, "+view)
	})
	// http://localhost:8080/star/123abc
	server.GET("/star/*abc", func(ctx *gin.Context) {
		view := ctx.Param("abc")
		ctx.String(http.StatusOK, "star是, "+view)
	})

	// POST
	server.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, login!")
	})

	//go func() {
	//	s1 := gin.Default()
	//	s1.Run(":8081")
	//}()
	server.Run(":8080")
}

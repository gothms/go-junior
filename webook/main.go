package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/web"
	"time"
)

func main() {
	//handler := &web.UserHandler{}
	handler := web.NewUserHandler()
	server := gin.Default()
	server.Use(func(ctx *gin.Context) {
		println("第一个 middleware")
	}, func(ctx *gin.Context) {
		println("第2个 middleware")
	})

	// middleware：跨域
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"}, // 允许的前端域名
		//AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是一个 middleware 示例")
	})

	handler.RegisterRoutes(server)
	server.Run(":8080")
}

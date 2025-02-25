package main

import (
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/web"
)

func main() {
	handler := &web.UserHandler{}
	server := gin.Default()
	handler.RegisterRoutes(server)

	server.Run(":8080")
}

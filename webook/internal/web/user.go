package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	// Restful 风格
	//server.POST("/user", h.SignUp)
	//server.PUT("/user", h.edit)
	//server.GET("/user/:username", h.Profile)

	//server.POST("/users/signup", h.SignUp)
	//server.POST("/users/login", h.Login)
	//server.POST("/users/edit", h.Edit)
	//server.GET("/users/profile", h.Profile) // 获取用户的基本信息

	// 分组
	ug := server.Group("/users")
	ug.POST("/signup", h.SignUp) // POST /users/signup
	ug.POST("/login", h.Login)
	ug.POST("/edit", h.Edit)
	ug.GET("/profile", h.Profile) // 获取用户的基本信息
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	ctx.String(http.StatusOK, "signup")
}

func (h *UserHandler) Login(ctx *gin.Context) {

}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {

}

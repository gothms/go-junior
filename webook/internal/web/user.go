package web

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	//server.POST("/users/sighup", h.SignUp)
	//server.POST("/users/login", h.Login)
	//server.POST("/users/edit", h.Edit)
	//server.POST("/users/profile", h.Profile)

	// 分组
	ug := server.Group("/users")
	ug.POST("/sighup", h.SignUp)
	ug.POST("/login", h.Login)
	ug.POST("/edit", h.Edit)
	ug.POST("/profile", h.Profile)
}
func (h *UserHandler) SignUp(ctx *gin.Context) {

}
func (h *UserHandler) Login(ctx *gin.Context) {

}
func (h *UserHandler) Edit(ctx *gin.Context) {

}
func (h *UserHandler) Profile(ctx *gin.Context) {

}

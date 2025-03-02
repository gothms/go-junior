package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleWareBuilder struct {
}

func (m *LoginMiddleWareBuilder) CheckLoginBuild() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			return // 注册和登录，不需要登录校验
		}
		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			// 中断，不要往后执行，即不执行后面业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

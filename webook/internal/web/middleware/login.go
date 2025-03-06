package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddleWareBuilder struct {
}

func (m *LoginMiddleWareBuilder) CheckLoginBuild() gin.HandlerFunc {
	gob.Register(time.Now()) // 注册一下 time.Time 类型
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			return // 注册和登录，不需要登录校验
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			// 中断，不要往后执行，即不执行后面业务逻辑
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()

		// 怎么知道要刷新登录状态了呢？
		// 假如说，策略是每分钟刷一次，如何知道过了一分钟？
		const updateTimeKey = "update_time"
		val := sess.Get(updateTimeKey) // 第一次进来
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || !ok || now.Sub(lastUpdateTime) > time.Second*10 {
			sess.Set(updateTimeKey, now)
			sess.Set("userId", id)
			err := sess.Save()
			if err != nil {
				fmt.Printf("user middleware:%v", err)
			}
		}
	}
}

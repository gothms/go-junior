package middleware

import (
	"encoding/gob"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-junior/webook/internal/web"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddleWareBuilder struct {
}

func (m *LoginJWTMiddleWareBuilder) CheckLogin() gin.HandlerFunc {
	gob.Register(time.Now()) // 注册一下 time.Time 类型
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		if path == "/users/signup" || path == "/users/login" {
			return // 注册和登录，不需要登录校验
		}
		// 根据约定，token 在 Authorization 头部
		authCode := ctx.GetHeader("authorization")
		if authCode == "" {
			// 没有登录，没有 token，Authorization 这个头部没有
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(authCode, " ")
		if len(segs) != 2 {
			// Authorization 是乱传的
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		var uc web.UserClaims
		token, err := jwt.ParseWithClaims(tokenStr, &uc, func(token *jwt.Token) (interface{}, error) {
			return web.JWTKey, nil
		})
		if err != nil {
			// token 不对，是伪造
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			// 非法 / 过期 token
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if uc.UserAgent != ctx.GetHeader("User-Agent") {
			// 当有监控告警时，这个地方要埋点
			// 能进来这个分支的，大概率是攻击者
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		expireTime := uc.ExpiresAt
		//if expireTime.Before(time.Now())	// token.Valid = false
		// 剩余过期时间 < 50s 就要刷新
		if d := expireTime.Sub(time.Now()); d >= 0 && d < time.Minute*4 {
			uc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
			tokenStr, err = token.SignedString(web.JWTKey)
			ctx.Header("x-jwt-token", tokenStr)
			if err != nil {
				// 不能中断，因为仅刷新过期时间失败
				fmt.Printf("user middleware:%v", err)
			}
		}
		ctx.Set("user", uc)
	}
}

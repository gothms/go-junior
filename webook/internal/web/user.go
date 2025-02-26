package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	emailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordRegexPattern = `^(?=.*\d)(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$`
)

type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		//emailRegexp:    regexp.MustCompile(emailRegexPattern),
		//passwordRegexp: regexp.MustCompile(passwordRegexPattern),
		emailRegexp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexp: regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
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

// SignUp 对应 webook-fe/src/pages/users/signup.tsx
// 正则表达式：
// 邮箱：^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$
// 密码（至少8位，必须含数字和特殊字符）：^(?=.*\d)(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$
func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignupReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignupReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//isEmail, err := regexp.Match(emailRegexPattern, []byte(req.Email))
	//if err != nil {
	//	ctx.String(http.StatusOK, "系统错误")
	//	return
	//}
	//isEmail := h.emailRegexp.Match([]byte(req.Email))
	isEmail, err := h.emailRegexp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		ctx.String(http.StatusOK, "非法邮箱")
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.String(http.StatusOK, "两次输入密码不同")
		return
	}

	isPassword, err := h.passwordRegexp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式不对，必须包含字母、数字、特殊字符，且不少于八位")
		return
	}

	ctx.String(http.StatusOK, "hello, 你在注册:"+req.Email)
}

func (h *UserHandler) Login(ctx *gin.Context) {

}

func (h *UserHandler) Edit(ctx *gin.Context) {

}

func (h *UserHandler) Profile(ctx *gin.Context) {

}

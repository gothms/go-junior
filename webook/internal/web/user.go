package web

import (
	"encoding/json"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/domain"
	"go-junior/webook/internal/service"
	"net/http"
	"time"
	"unicode/utf8"
)

const (
	emailRegexPattern    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	passwordRegexPattern = `^(?=.*\d)(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,72}$`
)

type UserHandler struct {
	emailRegexp    *regexp.Regexp
	passwordRegexp *regexp.Regexp
	svc            *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		//emailRegexp:    regexp.MustCompile(emailRegexPattern),
		//passwordRegexp: regexp.MustCompile(passwordRegexPattern),
		emailRegexp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordRegexp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		svc:            svc, // 依赖注入
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

	err = h.svc.Signup(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	// 判定邮箱冲突
	switch err {
	case nil:
		ctx.String(http.StatusOK, "hello, 你在注册:"+req.Email)
	case service.ErrDuplicateEmail:
		ctx.String(http.StatusOK, "邮箱冲突，请换一个")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	//fmt.Println(req.Email, req.Password)
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		sess := sessions.Default(ctx)
		sess.Set("userId", u.Id)
		sess.Options(sessions.Options{
			MaxAge: 900, // 15 分钟
		})
		err = sess.Save() // Gin 需要 Save
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
			return
		}

		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) Edit(ctx *gin.Context) {
	// nickname birthday personalProfile
	type EditReq struct {
		Nickname string `json:"nickname"`
		Birthday string `json:"birthday"`
		Personal string `json:"personal"`
	}
	var req EditReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	nickname, birthday, personal := req.Nickname, req.Birthday, req.Personal
	const (
		NicknameMaxLen        = 63
		BirthdayFormat        = time.DateOnly
		PersonalProfileMaxLen = 255
	)
	// check nickname
	if utf8.RuneCountInString(nickname) > NicknameMaxLen {
		ctx.String(http.StatusOK, "昵称最多有 %d 个字符（含中英文）", NicknameMaxLen)
		return
	}
	// check birthday
	parse, err := time.Parse(BirthdayFormat, birthday)
	if err != nil || parse.Format(BirthdayFormat) != birthday {
		//fmt.Println(err, parse, parse.Format(BirthdayFormat))
		ctx.String(http.StatusOK, "生日格式应为 YYYY-MM-DD，如：%s", BirthdayFormat)
		return
	}
	// check personal profile
	if len(personal) > PersonalProfileMaxLen {
		ctx.String(http.StatusOK, "个人简介最多有 %d 个字符", PersonalProfileMaxLen)
		return
	}

	id := h.getIdFromSession(ctx)
	err = h.svc.Edit(ctx, id, nickname, birthday, personal)
	switch err {
	case nil:
		ctx.String(http.StatusOK, "userId = %d，资料修改成功", id)
	default:
		// Error 1292 (22007): Truncated incorrect INTEGER value: ''
		ctx.String(http.StatusOK, "修改失败")
	}
	if err != nil {
		return
	}
}

func (h *UserHandler) Profile(ctx *gin.Context) {
	id := h.getIdFromSession(ctx)
	u, err := h.svc.Profile(ctx, id)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	jsonStr, err := json.Marshal(struct {
		Nickname string `json:"nickname"`
		Birthday string `json:"birthday"`
		Personal string `json:"personal"`
	}{
		Birthday: u.Birthday,
		Nickname: u.Nickname,
		Personal: u.Personal,
	})
	switch err {
	case nil:
		ctx.String(http.StatusOK, string(jsonStr))
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

func (h *UserHandler) getIdFromSession(ctx *gin.Context) int64 {
	return sessions.Default(ctx).Get("userId").(int64)
}

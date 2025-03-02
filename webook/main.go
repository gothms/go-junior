package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go-junior/webook/internal/repository"
	"go-junior/webook/internal/repository/dao"
	"go-junior/webook/internal/service"
	"go-junior/webook/internal/web"
	"go-junior/webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db := initDB()

	server := initWebServer()

	initUserHandler(db, server)

	server.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}
func initWebServer() *gin.Engine {
	server := gin.Default()

	//server.Use(func(ctx *gin.Context) {
	//	println("第一个 middleware")
	//}, func(ctx *gin.Context) {
	//	println("第2个 middleware")
	//})

	// middleware：跨域
	server.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"}, // 允许的前端域名
		//AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"authorization", "content-type"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是一个 middleware 示例")
	})

	login := middleware.LoginMiddleWareBuilder{}
	// 存储数据，也就是 userId 存的地方，这里直接存 Cookie
	store := cookie.NewStore([]byte("secret"))
	sess := sessions.Sessions("ssid", store) // 初始化 session
	server.Use(sess, login.CheckLoginBuild())

	return server
}
func initUserHandler(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)

	//handler := &web.UserHandler{}
	handler := web.NewUserHandler(us)

	handler.RegisterRoutes(server)
}

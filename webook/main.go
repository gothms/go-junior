package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go-junior/webook/internal/repository"
	"go-junior/webook/internal/repository/dao"
	"go-junior/webook/internal/service"
	"go-junior/webook/internal/web"
	"go-junior/webook/internal/web/middleware"
	"go-junior/webook/pkg/ginx/middleware/ratelimit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	//db := initDB()
	//
	//server := initWebServer()
	//
	//initUserHandler(db, server)
	server := gin.Default()
	server.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, 启动成功！")
	})

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
		//AllowHeaders: []string{"content-type"},
		AllowHeaders: []string{"authorization", "content-type"},
		// 允许前端访问后端响应中带的头部
		ExposeHeaders: []string{"x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是一个 middleware 示例")
	})

	// 限流插件
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	useJWT(server)

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
func useJWT(server *gin.Engine) {
	login := middleware.LoginJWTMiddleWareBuilder{}
	server.Use(login.CheckLogin())
}
func useSession(server *gin.Engine) {
	login := middleware.LoginMiddleWareBuilder{}
	// 存储数据，也就是 userId 存的地方
	// 方式一：直接存 Cookie
	store := cookie.NewStore([]byte("secret"))
	// 方式二：基于内存的实现
	// memstore：搜索生成32/64位的密钥
	//store := memstore.NewStore([]byte("MTc0NjY4MTY4NDIwMDY1NjAzNDI4MjIyNjgwNTAyMjYyNzY0MzAwODc3MjUwNjIwNTExMTc0ODczMzU3NDg0MQ=="),
	//	[]byte("MjY4NzQyMzYwNjM1NzYyMTQyNzc4MDEzNzU1NjMxNjc2NjQ3ODY2NTI2MTU2Mzc3NzMwMTg1MjQ1MTIyNzcxMA=="))
	// 方式三：redis
	// Authentication：是指身份认证，Encryption：是指数据加密
	//store, err := redis.NewStore(8, "tcp", "localhost:6379", "",
	//	[]byte("MTc0NjY4MTY4NDIwMDY1NjAzNDI4MjIyNjgwNTAyMjYyNzY0MzAwODc3MjUwNjIwNTExMTc0ODczMzU3NDg0MQ=="),
	//	[]byte("SJuumCxv4u3b83VEktwfY56wZ2szrEUc"))
	//if err != nil {
	//	panic(err)
	//}

	sess := sessions.Sessions("ssid", store) // 初始化 session
	server.Use(sess, login.CheckLoginBuild())
}

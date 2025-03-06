package bugs

/*
跨域
	CORS error：cors.jpg
		在 preflight 请求中告诉浏览器，允许 localhost:3000 发来的请求
	图示：preflight.jpg
		需要勾选 Disable cache，就能看到 preflight 请求
	gin：提供了 cors 解决方案
		https://github.com/gin-gonic/contrib：非常多 middleware
	解决：github.com/gin-contrib/cors
		https://github.com/gin-contrib/cors
	补充：修改 cors 代码后，可能不及时生效
		重启/重启电脑，根据 Network 的 preflight 重新设置参数
正则表达式
	panic: regexp: Compile(`^(?=.*\d)(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,}$`): error parsing regexp: invalid or unsupported Perl syntax: `(?=`
	解决：使用非官方的lib
		github.com/dlclark/regexp2
Gorm
	[error] failed to initialize database, got error dial tcp 127.0.0.1:3306: connec
	tex: No connection could be made because the target machine actively refused it.

	[error] failed to initialize database, got error dial tcp 127.0.0.1:3306: connec
		tex: No connection could be made because the target machine actively refused it.
	原因：mysql 服务未启动
		services.msc
	登录：
		mysql -u root -p -h 127.0.0.1
		mysql -u root -p
time.Time
	数据库：int64 / sql.NullInt64
	domain：time.Time
	string 转 time.Time：time.Parse(time.DateOnly, birthday)
	time.Time 转 string：timeVal.Format(time.DateOnly)
	time.Time 转 int64：timeVal.Unix()
	int64 转 time.Time：time.Unix(intVal, 0)
gob
	gob: type not registered for interface: time.Time
		session 写入 redis 中的类型为 Go 的结构体 time.Time，需要用到 gob 的序列化机制
		redis 中存入的是字节切片
	注册
		gob.Register(time.Now())
jwt & 跨域
	前端代码没有拿到 "x-jwt-token"，需要允许前端访问后端响应中带的头部
		AllowHeaders: []string{"authorization"},
		ExposeHeaders: []string{"x-jwt-token"},



*/
